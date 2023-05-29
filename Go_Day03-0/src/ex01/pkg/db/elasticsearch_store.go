package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ElasticsearchStore struct {
	esClient *elasticsearch.Client
	index    string
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

func (store *ElasticsearchStore) GetPlaces(limit, offset int) ([]Place, int, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	queryBytes, err := json.Marshal(query)
	if err != nil {
		return nil, 0, err
	}

	searchRequest := esapi.SearchRequest{
		Index:  []string{store.index},
		Body:   bytes.NewBuffer(queryBytes),
		Size:   &limit,
		Scroll: 1 * time.Minute,
	}

	searchResult, err := searchRequest.Do(context.Background(), store.esClient)
	if err != nil {
		return nil, 0, err
	}
	defer searchResult.Body.Close()

	if searchResult.IsError() {
		return nil, 0, fmt.Errorf("search error: %s", searchResult.String())
	}

	var response SearchResponse

	if err := json.NewDecoder(searchResult.Body).Decode(&response); err != nil {
		return nil, 0, err
	}

	totalHits := response.Hits.Total.Value

	places := make([]Place, len(response.Hits.Hits))
	for i, hit := range response.Hits.Hits {
		places[i] = hit.Source
	}

	return places, totalHits, nil
}
