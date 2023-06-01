package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func Example1() {
	// Создаем подключение к Elasticsearch
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"}, // Замените на соответствующий адрес и порт Elasticsearch
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Ошибка создания клиента Elasticsearch: %s", err)
	}

	// Проверяем соединение с Elasticsearch
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Ошибка получения информации о Elasticsearch: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Ошибка получения информации о Elasticsearch: %s", res.Status())
	}

	// Выводим информацию о версии Elasticsearch
	fmt.Println("Информация о версии Elasticsearch:")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(res.String())
	fmt.Println(strings.Repeat("=", 50))

	// Здесь вы можете добавить код для выполнения других операций с Elasticsearch

	// Пример создания нового документа
	createDocument(es)

	// Пример поиска документов
	searchDocuments(es)
}

func createDocument(es *elasticsearch.Client) {
	// Создаем JSON-документ
	doc := `{
		"title": "Пример документа",
		"description": "Это пример документа для Elasticsearch"
	}`

	// Создаем запрос для индексации документа
	req := esapi.IndexRequest{
		Index:      "my_index",
		DocumentID: "1",
		Body:       strings.NewReader(doc),
		Refresh:    "true",
	}

	// Выполняем запрос
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Ошибка при создании документа: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Ошибка при создании документа: %s", res.Status())
	}

	// Выводим результат
	fmt.Println("Документ успешно создан:")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(res.String())
	fmt.Println(strings.Repeat("=", 50))
}

func searchDocuments(es *elasticsearch.Client) {
	// Создаем запрос для поиска документов
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "пример",
			},
		},
	}
	queryBytes, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("Ошибка при сериализации запроса: %s", err)
	}
	req := esapi.SearchRequest{
		Index: []string{"my_index"},
		Body:  strings.NewReader(string(queryBytes)),
	}

	// Выполняем запрос
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Ошибка при поиске документов: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Ошибка при поиске документов: %s", res.Status())
	}

	// Выводим результат
	fmt.Println("Результаты поиска:")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println(res.String())
	fmt.Println(strings.Repeat("=", 50))
}
