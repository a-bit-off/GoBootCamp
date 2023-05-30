package dataBaseSite

import (
	"bytes"
	"encoding/json"
	"ex03/pkg/db"
	"log"
	"net/http"
)

type dataBase struct {
	Name   string     `json:"name"`
	Places []db.Place `json:"places"`
}

func DataBaseSite() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/recommend", handler)
	log.Fatal(http.ListenAndServe(":8888", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	places, err := db.New().GetPlaces(3, lat, lon)
	if err != nil {
		log.Fatal(err)
	}
	data := dataBase{Name: "Recommendation", Places: places}
	w.Header().Set("Content-Type", "application/json")
	marJSON, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Cannot marshal json", http.StatusInternalServerError)
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, marJSON, "", "\t")
	w.Write(prettyJSON.Bytes())
}
