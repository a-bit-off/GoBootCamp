package main

import (
	"context"
	"ex01/internal/application"
	"ex01/internal/repository"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var dbURL = "postgres://hazzeasu:password@localhost:5434/postgres" // ссылка на подключение к бд

func main() {
	// создаем контекст для безопасного завершения программы
	ctx := context.Background()

	// подключаемся к бд
	dbpool, err := repository.InitDbConnect(ctx, dbURL)
	if err != nil {
		log.Fatalf("%w failed to init DB connection", err)
	}
	defer dbpool.Close()

	// создаем приложение
	a := application.NewApp(ctx, dbpool)

	// создаем все таблицы если их нет
	err = a.CreateAllTables()
	if err != nil {
		log.Fatalf("%w failed to create tables", err)
	}

	// создаем роутер
	r := httprouter.New()
	a.Routes(r)

	// запускаем сервер
	srv := &http.Server{Addr: "localhost:8888", Handler: r}
	fmt.Println("Server listening")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("%w failed to listening server", err)
	}
	fmt.Println("server end")

}
