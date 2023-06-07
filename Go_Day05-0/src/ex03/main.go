package main

import (
	"ex03/pkg/combine"
	"ex03/pkg/myheap"
	"fmt"
)

func main() {
	presents := []myheap.Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}}
	for i := 0; i < 15; i++ {
		if res, err := grabPresents(presents, i); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Knapsack: %d\t->\t%d\n", i, res)
		}
	}
}

func grabPresents(presents []myheap.Present, maxSize int) ([]myheap.Present, error) {
	combinations := combine.GenerateCombinations(len(presents))
	var winCombination []int
	maxValue := 0

	for _, combination := range combinations {
		size, value := 0, 0
		for i := 0; i < len(combination); i++ {
			value += presents[combination[i]-1].Value
			size += presents[combination[i]-1].Size
		}
		if size <= maxSize {
			if maxValue < value {
				maxValue = value
				winCombination = combination
			}
		}
	}

	lWin := len(winCombination)
	res := make([]myheap.Present, lWin)
	for i := 0; i < lWin; i++ {
		res[i] = presents[winCombination[i]-1]
	}
	return res, nil
}
