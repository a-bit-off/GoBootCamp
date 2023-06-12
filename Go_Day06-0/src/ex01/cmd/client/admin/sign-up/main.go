package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AdminData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

var baseURL = "http://localhost:8888/admin/sign-up"

func main() {
	// create new Admin
	a1 := AdminData{"a9", "9"}

	// encode a1 to bytes
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(a1)

	// create new client
	client := http.Client{}

	// send request with new admin data and get response
	resp, err := client.Post(baseURL, "application/json", bytes.NewBuffer(data.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read response
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print response
	fmt.Println(string(rb))
}

// admin POST
// curl -XPOST -d '{"name": "p1", "password": "1"}' http://127.0.0.1:8888/admin
