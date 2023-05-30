package main

import (
	dbs "ex02/pkg/dataBaseSite"
	"ex02/pkg/db"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewElasticsearchStore(es, "places")

	places, totalHits, err := store.GetPlaces(10000, 0)
	if err != nil {
		log.Fatal(err)
	}

	dbs.DataBaseSite(places, totalHits)
}
