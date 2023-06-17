package application

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"ex01/internal/repository"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type app struct {
	ctx   context.Context
	repo  *repository.Repository
	cache map[string]repository.User
}

type postPage struct {
	Page    int
	Posts   []repository.Post
	Buttons buttons
}

type buttons struct {
	Previous int `json:"previous"`
	Next     int `json:"next"`
	Last     int `json:"last"`
}

var postsCountOnPage = 3 // колличество постов на странице
var allPostsCount int    // колличество всех постов

func NewApp(ctx context.Context, dbpool *pgxpool.Pool) *app {
	return &app{ctx, repository.NewRepository(dbpool), make(map[string]repository.User)}
}

func (a app) CreateAllTables() (err error) {
	err = a.repo.CreateTablePosts(a.ctx)
	if err != nil {
		return
	}

	a.repo.CreateTableUsers(a.ctx)
	if err != nil {
		return
	}

	allPostsCount, err = a.repo.GetAllPostsCount(a.ctx)
	if err != nil {
		return
	}

	return
}

func (a app) Routes(r *httprouter.Router) {
	r.ServeFiles("/../public/*filepath", http.Dir("public"))

	// startPage
	r.GET("/", a.StartPage)

	// newPost
	r.GET("/newPost", a.authorized(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		a.NewPostPage(w, "")
	}))
	r.POST("/newPost", a.authorized(a.NewPost))

	// login
	r.GET("/login", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		a.LoginPage(w, "")
	})
	r.POST("/login", a.Login)

	//logout
	r.GET("/logout", a.Logout)

	//signup
	r.GET("/signup", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		a.SignupPage(w, "")
	})
	r.POST("/signup", a.Signup)
}

func (a app) StartPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	posts, err := a.repo.GetNPosts(a.ctx, (page-1)*3)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	lp := filepath.Join("../public", "html", "startPage.html")
	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	last := allPostsCount / postsCountOnPage
	if allPostsCount%postsCountOnPage != 0 {
		last++
	}

	pp := postPage{page, posts,
		buttons{page - 1, page + 1, last}}
	err = tmpl.Execute(w, pp)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (a app) NewPostPage(w http.ResponseWriter, message string) {
	lp := filepath.Join("../public", "html", "newPost.html")
	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	type answer struct {
		Message string
	}

	data := answer{message}
	err = tmpl.ExecuteTemplate(w, "newPost", data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (a app) NewPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	header := r.FormValue("header")
	content := r.FormValue("content")

	if header == "" || content == "" {
		http.Error(w, "Пост должен содержать заголовок и содержание!", 400)
		return
	}

	token, err := readCookie("token", r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	user, ok := a.cache[token]
	if !ok {
		http.Error(w, fmt.Sprintf("Пользователь не определен!"), 400)
		return
	}
	err = a.repo.AddNewPost(a.ctx, user.Login, header, content)
	if err != nil {
		a.NewPostPage(w, fmt.Sprintf("Ошибка содания нового поста %s", err.Error()))
		return
	}
	allPostsCount++
	r.Form.Add("page", "1")
	a.StartPage(w, r, p)
}

func (a app) LoginPage(w http.ResponseWriter, message string) {
	lp := filepath.Join("../public", "html", "login.html")
	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	type answer struct {
		Message string
	}
	data := answer{message}
	err = tmpl.ExecuteTemplate(w, "login", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a app) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	login := r.FormValue("login")
	password := r.FormValue("password")
	if login == "" || password == "" {
		http.Error(w, fmt.Sprintf("Необходимо указать логин и пароль!"), 400)
		return
	}

	hash := md5.Sum([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])
	user, err := a.repo.Login(a.ctx, login, hashedPassword)
	if err != nil {
		a.LoginPage(w, fmt.Sprintf("Вы ввели неверный логин или пароль!: %s", err.Error()))
		return
	}

	// генерируем токен пишем его в кэш и cookie
	time64 := time.Now().Unix()
	timeInt := string(time64)
	token := login + password + timeInt
	hashToken := md5.Sum([]byte(token))
	hashedToken := hex.EncodeToString(hashToken[:])
	a.cache[hashedToken] = user
	livingTime := 60 * time.Minute
	expiration := time.Now().Add(livingTime)
	// кука будет жить 1 час
	cookie := http.Cookie{Name: "token", Value: url.QueryEscape(hashedToken), Expires: expiration}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/?page=1", http.StatusSeeOther)
}

func (a app) Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	for _, v := range r.Cookies() {
		c := http.Cookie{
			Name:   v.Name,
			MaxAge: -1}
		http.SetCookie(w, &c)
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (a app) SignupPage(w http.ResponseWriter, message string) {
	lp := filepath.Join("../public", "html", "signup.html")
	tmpl, err := template.ParseFiles(lp)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	type answer struct {
		Message string
	}
	data := answer{message}
	err = tmpl.ExecuteTemplate(w, "signup", data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func (a app) Signup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// получаем данные формы
	name := strings.TrimSpace(r.FormValue("name"))
	surname := strings.TrimSpace(r.FormValue("surname"))
	login := strings.TrimSpace(r.FormValue("login"))
	password := strings.TrimSpace(r.FormValue("password"))
	password2 := strings.TrimSpace(r.FormValue("password2"))

	// проверяем валидность данных
	if name == "" || surname == "" || login == "" || password == "" {
		a.SignupPage(w, "Все поля должны быть заполнены!")
		return
	}

	// сравниваем пароли
	if password != password2 {
		a.SignupPage(w, "Пароли не совпадают! Попробуйте еще")
		return
	}

	//	хешируем пароли
	hash := md5.Sum([]byte(password))
	hashedPassword := hex.EncodeToString(hash[:])

	// добавляем нового пользователя
	err := a.repo.AddNewUser(a.ctx, name, surname, login, hashedPassword)
	if err != nil {
		a.SignupPage(w, fmt.Sprintf("Ошибка создания пользователя %s", err.Error()))
		return
	}
	a.LoginPage(w, fmt.Sprintf("%s, Вы успешно зарегистрированы!", name))
}

func (a app) authorized(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token, err := readCookie("token", r)
		// если нет куки с токеном - пользователь отправляется на авторизацию
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		// если токен есть, но его нет в кеше - пользователь отправляется на авторизацию
		if _, ok := a.cache[token]; !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r, p)
	}
}

func readCookie(name string, r *http.Request) (value string, err error) {
	if name == "" {
		return value, errors.New("Вы пытаетесь прочитать пустой cookie")
	}
	cookie, err := r.Cookie(name)
	if err != nil {
		return
	}
	str := cookie.Value
	value, _ = url.QueryUnescape(str)
	return
}
