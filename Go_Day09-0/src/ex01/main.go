package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	in := make(chan string)
	go func() {
		defer close(in)
		in <- "https://a-bit-off.github.io/"
		in <- "https://a-bit-off.github.io/"
		in <- "https://a-bit-off.github.io/"
	}()

	result := crawlWeb(in)
	for res := range result {
		fmt.Println(res)
	}
}

func crawlWeb(input chan string) chan string {
	out := make(chan string, 8)
	go func() {
		defer close(out)
		var wg sync.WaitGroup
		defer wg.Wait()
		for in := range input {
			wg.Add(1)
			go func(in string) {
				defer wg.Done()
				parse, _ := parseHTML(in)
				out <- string(parse)
			}(in)
		}
	}()
	return out
}

func parseHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	parse, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(parse), nil
}
