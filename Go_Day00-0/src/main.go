package main

import (
	"fmt"
	"mean"
	"median"
	"mode"
	"os"
	p "parser"
	"sd"
)

type Metrics struct {
	mean   float64
	median float64
	mode   float64
	sd     float64
}

func main() {
	metrics := Metrics{}
	var option int = 1
	var read *os.File

	if option == 0 {
		if f, err := os.Open("./test/test.txt"); err == nil {
			read = f // read from file
		}
	} else if option == 1 {
		read = os.Stdin // read stdin
	}

	if order, err := p.ParseOrder(read); err == nil {
		fmt.Println("order:", order)
	} else {
		fmt.Println("err:", err)

	}

	if data, err := p.ParserData(read); err == nil {
		if metrics.mean, err = mean.Mean(data); err != nil {
			fmt.Println(err)
		}
		if metrics.median, err = median.Median(data); err != nil {
			fmt.Println(err)
		}
		if metrics.mode, err = mode.Mode(data); err != nil {
			fmt.Println(err)
		}
		if metrics.sd, err = sd.SD(data); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}

	fmt.Println(metrics)
}
