package server

import (
	"encoding/json"
	"ex01/server/pkg/db/users"
	"fmt"
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/admin", adminHandler())
	http.HandleFunc("/", rootHandler())

	log.Fatal(http.ListenAndServe(":8888", nil))
}

func rootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Привет, мир!")
	}
}

func adminHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser users.UserData

		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Ошибка при чтении JSON-запроса", http.StatusBadRequest)
			return
		}
		users.AddNewUser(newUser)
		fmt.Println(newUser)
		fmt.Fprintln(w, "Привет, админ!")
	}
}
