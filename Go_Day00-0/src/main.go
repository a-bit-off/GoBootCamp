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
	var read *os.File = os.Stdin
	fmt.Print("Metrics:\n\tMean - 1\n\tMedian - 2\n\tMode - 3\n\tSD - 4\n\tDefault - enter\nChoose your metrics: ")

	if order, err := p.ParseOrder(read); err == nil {
		if data, err := p.ParserData(read); err == nil {
			if metrics.mean, err = mean.Mean(data); err != nil {
				fmt.Println(err)
				return
			}
			if metrics.median, err = median.Median(data); err != nil {
				fmt.Println(err)
				return
			}
			if metrics.mode, err = mode.Mode(data); err != nil {
				fmt.Println(err)
				return
			}
			if metrics.sd, err = sd.SD(data); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println(err)
			return
		}
		PrintResult(order, metrics)
	}
}

func PrintResult(order []int, metrics Metrics) {
	for _, v := range order {
		switch v {
		case 1:
			fmt.Printf("Mean: %.2f\n", metrics.mean)
		case 2:
			fmt.Printf("Media: %.2f\n", metrics.median)
		case 3:
			fmt.Printf("Mode: %.2f\n", metrics.mode)
		case 4:
			fmt.Printf("SD: %.2f\n", metrics.sd)
		default:
			fmt.Printf("Metric №%d does not exist\n", v)
		}
	}
	if len(order) == 0 {
		fmt.Printf("Mean: %.2f\n", metrics.mean)
		fmt.Printf("Media: %.2f\n", metrics.median)
		fmt.Printf("Mode: %.2f\n", metrics.mode)
		fmt.Printf("SD: %.2f\n", metrics.sd)
	}
}
