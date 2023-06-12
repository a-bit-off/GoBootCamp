/*
sever
*/
package server

import (
	"encoding/json"
	"ex01/pkg/db/admin"
	"ex01/pkg/db/post"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Server() {
	PostsHandlersFunc()
	AdminHandlersFunc()
	http.HandleFunc("/", rootHandler())

	log.Fatal(http.ListenAndServe(":8888", nil))
}

// post
func PostsHandlersFunc() {
	http.HandleFunc("/post", viewPostHandler())
	http.HandleFunc("/post/new-post", newPostHandler())
}

// view post
func viewPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.FormValue("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}

		var newPost post.Post
		result, err := newPost.GetPosts()
		if err != nil {
			http.Error(w, "Ошибка при чтении bd", http.StatusBadRequest)
			return
		}

		lnResult := len(result)
		if page <= 0 || (lnResult-page*3) < -2 {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}
		start, end := 0, lnResult-(page*3)
		start = end + 3
		if end < 0 {
			end = 0
		}

		for i := start - 1; i >= end; i-- {
			w.Write([]byte(fmt.Sprintf("\tPost№ %d\n\n", i)))
			w.Write([]byte(fmt.Sprintf("\tHeader: %s\n", result[i].Header)))
			w.Write([]byte(fmt.Sprintf("\tContent:\n%s\n", result[i].Content)))
			w.Write([]byte(fmt.Sprintf("\n\n\n")))
		}
	}
}

// new post
func newPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newPost post.Post

		err := json.NewDecoder(r.Body).Decode(&newPost)
		if err != nil {
			http.Error(w, "Ошибка при чтении JSON-запроса", http.StatusBadRequest)
			return
		}
		err = newPost.NewPost()
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при создании поста: %s", err), http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, "Пост успешно создан!")
	}
}

// admin
func AdminHandlersFunc() {
	http.HandleFunc("/admin/sign-up", adminSignUpHandler())
	http.HandleFunc("/admin/sign-in", adminSignInHandler())
}

// регистрация
func adminSignUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAdmin admin.AdminData

		err := json.NewDecoder(r.Body).Decode(&newAdmin)
		if err != nil {
			http.Error(w, "Ошибка при чтении JSON-запроса", http.StatusBadRequest)
			return
		}
		err = newAdmin.SignUpAdmin()
		if err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, "Привет, админ!")
	}
}

// вход
func adminSignInHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAdmin admin.AdminData

		err := json.NewDecoder(r.Body).Decode(&newAdmin)
		if err != nil {
			http.Error(w, "Ошибка при чтении JSON-запроса", http.StatusBadRequest)
			return
		}
		err = newAdmin.SignInAdmin()
		if err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
			return
		}
		fmt.Println(newAdmin)
		fmt.Fprintln(w, "С возвращением, админ!")
	}
}

// default
func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Привет!")
	}
}
