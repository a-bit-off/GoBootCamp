/*
Пакет myES считывает данные из CSV-файла,
создает клиент Elasticsearch с настройками по умолчанию,
создает индекс с необходимой схемой
и загружает данные в индекс

На вход принимает:

	dataSetPath: путь к csv файлу
	indexName: название индекса
	dataCount: количество данных необходимых для загрузки
*/
package myES

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func MyES(dataSetPath, indexName string, dataCount int) error {
	dataSet, err := readDataSet(dataSetPath)
	if err != nil {
		return err
	}

	// Установливаем соединение с сервером Elasticsearch.
	/*
		Натройки по умолчанию включают:
		1. Адрес сервера Elasticsearch: По умолчанию клиент
		настроен на подключение к серверу Elasticsearch,
		работающему на локальной машине (localhost:9200).
		2. Режим подключения: Клиент устанавливает соединение
		с сервером Elasticsearch в режиме HTTP.
		3. Таймауты: Установлены значения таймаутов для запросов
		к серверу Elasticsearch. По умолчанию установлен таймаут
		в 1 минуту для большинства операций.
		4. Журналирование: Клиент выводит журнальные сообщения,
		которые могут быть полезными для отладки и мониторинга.
		5. Ретрисы: Клиент автоматически выполняет повторные попытки
		при определенных ошибках, таких как проблемы с подключением
		или таймауты.
	*/
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return err
	}

	if err := createIndexWithMapping(es, indexName); err != nil {
		return err
	}

	if dataCount == 0 {
		dataCount = len(dataSet)
	}

	if err := uploadDataToIndex(es, indexName, dataSet, dataCount); err != nil {
		return err
	}
	return nil
}

// Читает с CSV-файла.
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

// Создает индекс Elasticsearch с соответствующей схемой.
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

// Загружает данные в индекс Elasticsearch.
func uploadDataToIndex(es *elasticsearch.Client, indexName string, dataSet [][]string, dataCount int) error {
	for i := 0; i < dataCount; i++ {
		place := createPlace(dataSet[i])
		jsonData, err := json.Marshal(place)
		if err != nil {
			return err
		}
		// запрос на индексацию в Elasticsearch
		request := esapi.IndexRequest{
			Index:        indexName,
			DocumentID:   strconv.Itoa(i),
			DocumentType: "place",
			Body:         bytes.NewReader(jsonData),
			Refresh:      "true",
		}

		// выполняем запрос на индексацию
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

// Создает объект Place на основе данных из набора данных.
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
