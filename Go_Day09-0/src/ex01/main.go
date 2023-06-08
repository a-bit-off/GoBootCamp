package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// part 2
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	go func() {
		sig := <-sigs
		fmt.Println("\nsignal:", sig)
		cancel()
	}()

	// part 1
	in := make(chan string)
	go func() {
		defer close(in)
		for i := 0; i < 1000; i++ {
			in <- "https://a-bit-off.github.io/"
		}
	}()

	result := crawlWeb(ctx, in)

	rightToFile(result)
}

func crawlWeb(ctx context.Context, input chan string) chan string {
	out := make(chan string, 8)

	select {
	case <-ctx.Done():
		defer close(out)
		return out

	default:
		go func() {
			defer close(out)
			var wg sync.WaitGroup
			defer wg.Wait()
		Loop:
			for in := range input {
				select {
				case <-ctx.Done():
					break Loop
				default:
					wg.Add(1)
					go func(in string) {
						defer wg.Done()
						parse, _ := parseHTML(in)
						out <- string(parse)
					}(in)
				}
			}
		}()
	}

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

func rightToFile(result <-chan string) {
	file, err := os.Create("parse.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for res := range result {
		file.WriteString(res)
	}
}
