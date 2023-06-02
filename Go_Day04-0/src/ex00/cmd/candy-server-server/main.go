package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"ex00/restapi"
	"ex00/restapi/operations"
)

type CandyRequest struct {
	Money      int64  `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int64  `json:"candyCount"`
}

type Candy struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

var candyPriceList = []Candy{
	{
		Name:  "CE",
		Price: 10,
	},
	{
		Name:  "AA",
		Price: 15,
	},
	{
		Name:  "NT",
		Price: 17,
	},
	{
		Name:  "DE",
		Price: 21,
	},
	{
		Name:  "YR",
		Price: 23,
	},
}

func main() {
	// Загрузка встроенной спецификации Swagger с помощью функции loads.Embedded()
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	// Создание экземпляра API сервера CandyServerAPI на основе загруженной спецификации Swagger.
	// Затем создается экземпляр сервера NewServer с использованием созданного API.
	// Отложенный вызов server.Shutdown() гарантирует, что сервер будет корректно остановлен при выходе из функции main
	api := operations.NewCandyServerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = "127.0.0.1"
	server.Port = 3333

	// обработка request
	api.BuyCandyHandler = operations.BuyCandyHandlerFunc(
		func(params operations.BuyCandyParams) middleware.Responder {
			// request
			candyRequest := CandyRequest{
				Money:      *params.Order.Money,
				CandyType:  *params.Order.CandyType,
				CandyCount: *params.Order.CandyCount,
			}

			// responses:
			// response 400
			candyPL, err := GetCandyFromPL(candyRequest)
			if err != nil {
				return operations.NewBuyCandyBadRequest().WithPayload(
					&operations.BuyCandyBadRequestBody{Error: fmt.Sprintf("%s", err)})
			}

			change := candyRequest.Money - (candyPL.Price * candyRequest.CandyCount)

			// response 402
			if change < 0 {
				return operations.NewBuyCandyPaymentRequired().WithPayload(
					&operations.BuyCandyPaymentRequiredBody{Error: fmt.Sprintf("Incorrect input: You need %d more money!", -change)})
			}

			// response 201
			return operations.NewBuyCandyCreated().WithPayload(
				&operations.BuyCandyCreatedBody{
					Change: change,
					Thanks: fmt.Sprintf("Thank you!"),
				})
		})

	// Запуск сервера на обслуживание входящих запросов методом Serve()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func RequestValid(candyReq CandyRequest) error {
	success := false
	for _, c := range candyPriceList {
		if c.Name == candyReq.CandyType {
			success = true
			break
		}
	}
	if !success || candyReq.CandyCount < 0 || candyReq.Money < 0 {
		return errors.New("Incorrect input: some error in input data")
	}
	return nil
}

// Get candy from price list
func GetCandyFromPL(candyReq CandyRequest) (Candy, error) {
	var findCandy Candy
	if err := RequestValid(candyReq); err != nil {
		return Candy{}, err
	}
	for _, c := range candyPriceList {
		if c.Name == candyReq.CandyType {
			findCandy = c
			break
		}
	}
	return findCandy, nil
}
