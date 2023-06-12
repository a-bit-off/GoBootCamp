package server

import (
	"encoding/json"
	"ex01/pkg/db/admins"
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
		var newAdmin admins.AdminData

		err := json.NewDecoder(r.Body).Decode(&newAdmin)
		if err != nil {
			http.Error(w, "Ошибка при чтении JSON-запроса", http.StatusBadRequest)
			return
		}
		admins.AddNewAdmin(newAdmin)
		fmt.Println(newAdmin)
		fmt.Fprintln(w, "Привет, админ!")
	}
}
