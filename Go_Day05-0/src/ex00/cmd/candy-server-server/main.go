// git solution
/*
package main

import (
	"errors"
	"ex00/restapi"
	"ex00/restapi/operations"
	"flag"
	"fmt"
	"log"
	"math"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

var (
	portFlag   = flag.Int("port", 3333, "Port to run this service on")
	candyPrice = make(map[string]int64)
)

type good struct {
	name  string
	price int64
}

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCandyServerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	flag.Parse()
	server.Port = *portFlag
	priceList := getPriceList()
	fillMap(priceList)
	api.BuyCandyHandler = operations.BuyCandyHandlerFunc(
		func(params operations.BuyCandyParams) middleware.Responder {
			candy := struct {
				name  string
				count int64
				money int64
			}{
				swag.StringValue(params.Order.CandyType),
				swag.Int64Value(params.Order.CandyCount),
				swag.Int64Value(params.Order.Money),
			}
			_, ok := candyPrice[candy.name]
			if !ok {
				return operations.NewBuyCandyBadRequest().WithPayload(
					&operations.BuyCandyBadRequestBody{Error: fmt.Sprintf("Incorrect input: we don't have %s", candy.name)})
			}

			if candy.count <= 0 || candy.money <= 0 {
				return operations.NewBuyCandyBadRequest().WithPayload(
					&operations.BuyCandyBadRequestBody{Error: "Incorrect input: money or count can't be zero or negative"})
			}

			change, err := calculateChange(candy.count, candyPrice[candy.name], candy.money)
			if err != nil {
				return operations.NewBuyCandyPaymentRequired().WithPayload(
					&operations.BuyCandyPaymentRequiredBody{Error: err.Error()})
			}

			return operations.NewBuyCandyCreated().WithPayload(
				&operations.BuyCandyCreatedBody{Change: change, Thanks: "Thank you!"})
		})

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func getPriceList() []good {
	return []good{
		{name: "CE", price: 10},
		{name: "AA", price: 15},
		{name: "NT", price: 17},
		{name: "DE", price: 21},
		{name: "YR", price: 23},
	}
}

func fillMap(priceList []good) {
	for _, v := range priceList {
		candyPrice[v.name] = v.price
	}
}

func calculateChange(count, price, money int64) (int64, error) {
	change := money - (price * count)
	if money < price*count {

		return 0, errors.New(fmt.Sprintf("You need %d more money!", int(math.Abs(float64(change)))))
	}
	return change, nil
}
*/

// chat solution
/*
package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"ex00/restapi"
	"ex00/restapi/operations"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewCandyServerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Candy Server"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.Port = 3333
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
*/

// my solution
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

type CandyResponse struct {
	Change int64  `json:"change"`
	Thanks string `json:"thanks"`
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
					&operations.BuyCandyPaymentRequiredBody{Error: fmt.Sprintf("Incorrect input: not enough money")})
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
