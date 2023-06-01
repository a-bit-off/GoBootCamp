package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
