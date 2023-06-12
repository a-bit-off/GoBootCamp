/*
view post
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	Header  string `json:"header"`
	Content string `json:"content"`
}

var baseURL = "http://localhost:8888/post"

func main() {
	// create new Post
	post := Post{"header", "content"}

	// encode a1 to bytes
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(post)

	// create new client
	client := http.Client{}

	// send request with new post data and get response
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
