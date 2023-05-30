package dataBaseSite

import (
	"ex03/pkg/db"
	"fmt"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
)

type dataBase struct {
	Name   string     `json:"name"`
	Places []db.Place `json:"places"`
}

func DataBaseSite() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewElasticsearchStore(es, "places")

	mux := http.NewServeMux()
	mux.HandleFunc("/api/", handler(store))
	log.Fatal(http.ListenAndServe(":8888", mux))
}

func handler(store *db.ElasticsearchStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lat := r.URL.Query().Get("lat")
		lon := r.URL.Query().Get("lon")
		fmt.Println("lat", lat)
		fmt.Println("lon", lon)
		places, err := store.GetPlaces(3, lat, lon)
		if err != nil {
			log.Fatal(err)
		}
		for _, place := range places {
			fmt.Println(place)
		}
		// data := dataBase{Name: "Recommendation", Places: places}
		// w.Header().Set("Content-Type", "application/json")
		// marJSON, err := json.Marshal(data)
		// if err != nil {
		// 	http.Error(w, "Cannot marshal json", http.StatusInternalServerError)
		// }

		// var prettyJSON bytes.Buffer
		// err = json.Indent(&prettyJSON, marJSON, "", "\t")
		// w.Write(prettyJSON.Bytes())
	}
}
