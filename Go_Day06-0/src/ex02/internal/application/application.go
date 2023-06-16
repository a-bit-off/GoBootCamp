package application

import (
	"context"
	"ex02/internal/repository"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
)

type app struct {
	ctx   context.Context
	repo  *repository.Repository
	cache map[string]repository.User
}

// колличество постов на странице
var postsCountOnPage = 3

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
	return
}

func (a app) Routes(r *httprouter.Router) {
	//r.ServeFiles() TODO в случае внешних соединений
	r.GET("/", a.StartPage)

	// post
	r.GET("/newPost", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		a.NewPostPage(rw, "")
	})
	r.POST("/newPost", a.NewPost)
}

// TODO довести до ума
func (a app) StartPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	posts, err := a.repo.GetNPosts(a.ctx, page*3, postsCountOnPage)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	var res string
	for _, p := range posts {
		res += fmt.Sprintf("userID: %d\nheader: %s\ncontnet: %s\n\tcreated: %s\n\n\n",
			p.UserId, p.Header, p.Content, p.Created)
	}
	w.Write([]byte(res))
}

func (a app) NewPostPage(w http.ResponseWriter, message string) {
	lp := filepath.Join("../public", "html", "newPost.html.go")
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
	}

	// TODO userId сделать через cach
	err := a.repo.AddNewPost(a.ctx, 0, header, content)
	if err != nil {
		a.NewPostPage(w, fmt.Sprintf("Ошибка содания нового поста %s", err.Error()))
		return
	}
}
