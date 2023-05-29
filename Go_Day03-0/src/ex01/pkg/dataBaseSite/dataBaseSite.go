package dataBaseSite

import (
	"ex01/pkg/db"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dataBase struct {
	Total   int        `json:"total"`
	Places  []db.Place `json:"places"`
	Buttons buttons    `json:"button"`
	Page    int        `json:"page"`
}

type buttons struct {
	Previous int `json:"previous"`
	Next     int `json:"next"`
	Last     int `json:"last"`
}

func DataBaseSite(places []db.Place, totalHits int, filePath string) {
	tmpl := template.Must(template.ParseFiles(filePath))
	http.HandleFunc("/", handler(places, totalHits, tmpl))
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func handler(places []db.Place, totalHits int, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.FormValue("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}

		start := page * 10
		end := (page + 1) * 10
		if start >= totalHits || start < 0 {
			http.Error(w, fmt.Sprintf("Page \"%d\" not found", page), http.StatusNotFound)
			return
		}
		if end > totalHits {
			end = totalHits
		}
		but := buttons{
			Previous: page - 1,
			Next:     page + 1,
			Last:     totalHits % 10,
		}
		data := dataBase{Total: totalHits, Places: places[start:end], Buttons: but, Page: page}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Page \"%d\" not found", page), http.StatusNotFound)
		}
	}
}
