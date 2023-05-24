package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

type Data struct {
	l, m, w bool
	paths   []string
}

func main() {
	data, err := Parse()
	if err != nil {
		fmt.Println(err)
		return
	} else if err := FlagCompatibility(data); err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	res := make([]string, 0)
	for i := 0; i < len(data.paths); i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			MyWC(data, k, &mu, &res)
		}(i)
	}
	wg.Wait()
	for _, r := range res {
		fmt.Println(r)
	}
}

func Parse() (Data, error) {
	l := flag.Bool("l", false, "l flag")
	m := flag.Bool("m", false, "m flag")
	w := flag.Bool("w", false, "w flag")

	flag.Parse()
	if len(os.Args) < 3 {
		return Data{}, errors.New("Missing flag or file path")
	}
	return Data{l: *l, m: *m, w: *w, paths: os.Args[2:]}, nil
}

func FlagCompatibility(data Data) error {
	flags := []bool{data.l, data.m, data.w}
	count := 0
	for _, f := range flags {
		if f {
			count++
		}
	}
	if count > 1 {
		return errors.New("Too many flags")
	}
	return nil
}

func MyWC(data Data, fileNameIndex int, mu *sync.Mutex, res *[]string) {
	file, err := os.Open(data.paths[fileNameIndex])
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	count := 0
	in := bufio.NewScanner(file)
	for in.Scan() {
		if err := in.Err(); err != nil {
			fmt.Println(err)
		}
		if data.l {
			count++
		} else if data.m {
			count += utf8.RuneCountInString(in.Text()) + 1
		} else {
			count += len(strings.Split(in.Text(), " "))
		}
	}
	if count != 0 && (data.l || data.m) {
		count--
	}

	mu.Lock()
	*res = append(*res, fmt.Sprintf("%d\t%s", count, data.paths[fileNameIndex]))
	mu.Unlock()
}
