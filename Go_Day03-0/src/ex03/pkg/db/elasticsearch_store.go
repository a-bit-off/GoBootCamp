package db

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/olivere/elastic"
)

type ElasticsearchStore struct {
	esClient *elasticsearch.Client
	index    string
}

type Place struct {
	Name     string           `json:"name"`
	Address  string           `json:"address"`
	Phone    string           `json:"phone"`
	Location elastic.GeoPoint `json:"location"`
}

type SearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Source Place `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func NewElasticsearchStore(esClient *elasticsearch.Client, index string) *ElasticsearchStore {
	return &ElasticsearchStore{
		esClient: esClient,
		index:    index,
	}
}

func (store *ElasticsearchStore) GetPlaces(limit int, latStr string, lonStr string) ([]Place, error) {
	if latStr == "" || lonStr == "" {
		return nil, errors.New("Empty lat / lon")
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return nil, err
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return nil, err
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"geo_distance": map[string]interface{}{
				"distance": "1000m", // Радиус поиска
				"location": map[string]interface{}{
					"lat": lat,
					"lon": lon,
				},
			},
		},
		"size": limit, // Ограничение количества результатов
	}

	queryBytes, err := json.Marshal(query)

	if err != nil {
		return nil, err
	}

	searchRequest := esapi.SearchRequest{
		Index:  []string{store.index},
		Body:   bytes.NewBuffer(queryBytes),
		Scroll: 1 * time.Minute,
	}

	searchResult, err := searchRequest.Do(context.Background(), store.esClient)

	if err != nil {
		return nil, err
	}

	defer searchResult.Body.Close()

	if searchResult.IsError() {
		return nil, fmt.Errorf("search error: %s", searchResult.String())
	}

	var response SearchResponse

	if err := json.NewDecoder(searchResult.Body).Decode(&response); err != nil {
		return nil, err
	}

	places := make([]Place, len(response.Hits.Hits))

	for i, hit := range response.Hits.Hits {
		places[i] = hit.Source
	}

	return places, nil
}
