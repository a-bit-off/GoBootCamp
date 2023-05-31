package main

import (
	"ex00/pkg/server"
)

func main() {
	server.Server()
}

// написать запрос либо вызвать отельно в терминале

// curl -XPOST -H "Content-Type: application/json" -d '{"money": 20, "candyType": "AA", "candyCount": 1}' http://127.0.0.1:3333/buy_candy
// curl -XPOST -H "Content-Type: application/json" -d '{"money": 46, "candyType": "YR", "candyCount": 2}' http://127.0.0.1:3333/buy_candy
