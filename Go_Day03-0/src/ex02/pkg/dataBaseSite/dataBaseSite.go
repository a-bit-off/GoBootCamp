package dataBaseSite

import (
	"bytes"
	"encoding/json"
	"ex02/pkg/db"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dataBase struct {
	Name     string     `json:"name"`
	Total    int        `json:"total"`
	Places   []db.Place `json:"places"`
	PrevPage int        `json"prev_page"`
	NextPage int        `json"next_page"`
	LastPage int        `json"last_page"`
}

func DataBaseSite(places []db.Place, totalHits int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", handler(places, totalHits))
	log.Fatal(http.ListenAndServe(":8888", mux))
}

func handler(places []db.Place, totalHits int) http.HandlerFunc {
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
		var prevP, nextP int
		if page-1 >= 0 {
			prevP = page - 1
		}
		if page+1 <= (totalHits / 10) {
			nextP = page + 1
		}
		data := dataBase{Name: "places", Total: totalHits, Places: places[start:end],
			PrevPage: prevP, NextPage: nextP, LastPage: totalHits / 10}
		w.Header().Set("Content-Type", "application/json")
		marJSON, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Cannot marshal json", http.StatusInternalServerError)
		}

		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, marJSON, "", "\t")
		w.Write(prettyJSON.Bytes())
	}
}
