package server

import (
	"encoding/json"
	"ex01/pkg/db/admins"
	"fmt"
	"log"
	"net/http"
)

func Server() {
	PostsHandlersFunc()
	AdminHandlersFunc()
	http.HandleFunc("/", rootHandler())

	log.Fatal(http.ListenAndServe(":8888", nil))
}

// handlesFunc
func PostsHandlersFunc() {
	http.HandleFunc("/post", postHandler())
}

func AdminHandlersFunc() {
	http.HandleFunc("/admin/sign-up", adminSignUpHandler())
	http.HandleFunc("/admin/sign-in", adminSignInHandler())
}

// post
func postHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

// admin
// регистрация
func adminSignUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAdmin admins.AdminData

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
		fmt.Println(newAdmin)
		fmt.Fprintln(w, "Привет, админ!")
	}
}

// вход
func adminSignInHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newAdmin admins.AdminData

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

// дефолт
func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Привет!")
	}
}
