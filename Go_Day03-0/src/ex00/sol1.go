package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/grailbio/base/tsv"
	"github.com/olivere/elastic"
)

type Place struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Address  string           `json:"address"`
	Phone    string           `json:"phone"`
	Location elastic.GeoPoint `json:"location"`
}

type IndexMapping struct {
	Properties MappingProperties `json:"properties"`
}

type MappingProperties struct {
	Name     FieldType `json:"name"`
	Address  FieldType `json:"address"`
	Phone    FieldType `json:"phone"`
	Location FieldType `json:"location"`
}

type FieldType struct {
	Type string `json:"type"`
}

func main() {
	dataSetPath := flag.String("dataSetPath", "../../materials/data.csv", "path to csv file")
	indexName := flag.String("indexName", "places", "index name")
	dataCount := flag.Int("dataCount", 0, "number of data entries to upload")

	flag.Parse()

	dataSet, err := readDataSet(*dataSetPath)
	if err != nil {
		log.Fatal(err)
	}

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	if err := createIndexWithMapping(es, *indexName); err != nil {
		log.Fatal(err)
	}
	fmt.Println("*dataCount:", *dataCount)
	if *dataCount == 0 {
		*dataCount = len(dataSet)
	}
	fmt.Println("*dataCount:", *dataCount)

	if err := uploadDataToIndex(es, *indexName, dataSet, *dataCount); err != nil {
		log.Fatal(err)
	}
}

func readDataSet(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := tsv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func createIndexWithMapping(es *elasticsearch.Client, indexName string) error {
	var buf bytes.Buffer
	mapping := IndexMapping{
		Properties: MappingProperties{
			Name:     FieldType{Type: "text"},
			Address:  FieldType{Type: "text"},
			Phone:    FieldType{Type: "text"},
			Location: FieldType{Type: "geo_point"},
		},
	}

	if err := json.NewEncoder(&buf).Encode(mapping); err != nil {
		return err
	}

	res, err := es.Indices.PutMapping(
		strings.NewReader(buf.String()),
		es.Indices.PutMapping.WithIndex(indexName),
		es.Indices.PutMapping.WithIncludeTypeName(true),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func uploadDataToIndex(es *elasticsearch.Client, indexName string, dataSet [][]string, dataCount int) error {
	for i := 0; i < dataCount; i++ {
		place := createPlace(dataSet[i])
		jsonData, err := json.Marshal(place)
		if err != nil {
			return err
		}

		request := esapi.IndexRequest{
			Index:        indexName,
			DocumentID:   strconv.Itoa(i + 1),
			DocumentType: "place",
			Body:         bytes.NewReader(jsonData),
			Refresh:      "true",
		}

		response, err := request.Do(context.Background(), es)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.IsError() {
			return fmt.Errorf("error indexing document: %s", response.Status())
		}

		fmt.Println("Status:", response.Status())
	}

	return nil
}

func createPlace(data []string) Place {
	id, _ := strconv.Atoi(data[0])
	lon, _ := strconv.ParseFloat(data[4], 64)
	lat, _ := strconv.ParseFloat(data[5], 64)
	return Place{
		ID:       id,
		Name:     data[1],
		Address:  data[2],
		Phone:    data[3],
		Location: elastic.GeoPoint{Lon: lon, Lat: lat},
	}
}
