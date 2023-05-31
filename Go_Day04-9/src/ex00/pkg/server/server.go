package server

import (
	"encoding/json"
	"errors"
	"ex00/pkg/candy"
	"fmt"
	"log"
	"net/http"
)

var candyPriceList = []candy.Candy{
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

func Server() {
	http.HandleFunc("/buy_candy", handler)
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Ожидается метод POST", http.StatusMethodNotAllowed)
		return
	}

	// Чтение JSON запроса
	var candyReq candy.CandyRequest
	err := json.NewDecoder(r.Body).Decode(&candyReq)
	if err != nil {
		http.Error(w, "Ошибка при чтении JSON-запроса", http.StatusBadRequest)
		return
	}

	// Обработка запроса
	if err = RequestValid(candyReq); err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
		return
	}

	candyData, err := FindCandyInPL(candyReq)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
		return
	}

	candyResp := candy.CandyResponse{
		Change: candyReq.Money - (candyReq.CandyCount * candyData.Price),
		Thanks: "Thank you!",
	}
	if candyResp.Change < 0 {
		candyResp.Thanks = fmt.Sprintf("You need %d more money!", -candyResp.Change)
		candyResp.Change = 0
	}

	// Кодирование JSON-ответа
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(candyResp)
	if err != nil {
		http.Error(w, "Ошибка при кодировании JSON-ответа", http.StatusInternalServerError)
		return
	}
}

func RequestValid(candyReq candy.CandyRequest) error {
	success := false
	for _, c := range candyPriceList {
		if c.Name == candyReq.CandyType {
			success = true
			break
		}
	}
	if !success || candyReq.CandyCount < 0 || candyReq.Money < 0 {
		return errors.New("Некорректный запрос")
	}
	return nil
}

func FindCandyInPL(candyReq candy.CandyRequest) (candy.Candy, error) {
	for _, c := range candyPriceList {
		if c.Name == candyReq.CandyType {
			return c, nil
		}
	}
	return candy.Candy{}, errors.New("\"CandyType\" не определен")
}
