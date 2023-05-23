package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	l, m, w bool
	paths   []string
}

func main() {
	data, err := Parse()
	if err != nil {
		fmt.Println(err)
	} else if err := FlagCompatibility(data); err != nil {
		fmt.Println(err)
	} else {
		channel := make(chan int)
		go WordCountGorutine(data, channel)
		fileCount := 0
		for ch := range channel {
			fmt.Println(ch, "\t", data.paths[fileCount])
			fileCount++
		}
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

func WordCountGorutine(data Data, channel chan int) {
	for i := 0; i < len(data.paths); i++ {
		WordCount(data, i, channel)
	}
	close(channel)
}

func WordCount(data Data, fileNameIndex int, channel chan int) {
	// open file
	file, err := os.Open(data.paths[fileNameIndex])
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// scan file
	var count int
	in := bufio.NewScanner(file)
	for in.Scan() {
		if err := in.Err(); err != nil {
			fmt.Println(err)
		}
		// count flag
		if data.l {
			count++
		} else if data.m {
			count += len(in.Text())
		} else {
			count += len(strings.Split(in.Text(), " "))
		}
	}
	channel <- count
}
