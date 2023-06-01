package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

type Place struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type Store interface {
	GetPlaces(limit, offset int) ([]Place, int, error)
}

type ElasticsearchStore struct {
	esClient *elasticsearch.Client
	index    string
}

func NewElasticsearchStore(esClient *elasticsearch.Client, index string) *ElasticsearchStore {
	return &ElasticsearchStore{
		esClient: esClient,
		index:    index,
	}
}

func (store *ElasticsearchStore) GetPlaces(limit, offset int) ([]Place, int, error) {
	query := map[string]interface{}{
		"size": limit,
		"from": offset,
	}

	body, err := json.Marshal(query)
	if err != nil {
		return nil, 0, err
	}

	response, err := store.esClient.Search(
		store.esClient.Search.WithIndex(store.index),
		store.esClient.Search.WithBody(bytes.NewBuffer(body)),
		store.esClient.Search.WithTrackTotalHits(true),
		store.esClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, 0, err
	}
	defer response.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, 0, err
	}

	totalHits := int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	places := make([]Place, 0)
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		id := hit.(map[string]interface{})["_id"].(string)
		place := Place{
			ID:      id,
			Name:    source["name"].(string),
			Address: source["address"].(string),
			Phone:   source["phone"].(string),
		}
		places = append(places, place)
	}

	return places, totalHits, nil
}

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	store := NewElasticsearchStore(es, "places")

	places, totalHits, err := store.GetPlaces(10, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Hits: %d\n", totalHits)
	for _, place := range places {
		fmt.Printf("ID: %s, Name: %s, Address: %s, Phone: %s\n", place.ID, place.Name, place.Address, place.Phone)
	}
}
