package main

import (
	"ex03/myheap"
	"fmt"
)

func main() {
	presents := []myheap.Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}}
	if res, err := grabPresents(presents, 6); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func grabPresents(presents []myheap.Present, maxSize int) ([]myheap.Present, error) {
	lenPres := len(presents)
	matrix := make([][]int, 0, lenPres)
	// цикл по количеству предметов
	for i := 0; i < lenPres; i++ {
		present := presents[i]
		mat := make([]int, 0, maxSize)
		// цикл по размеру рюкзака
		for j := 0; j < maxSize; j++ {
			// если пред помещается то мы кладем его
			if j >= present.Size {
				mat = append(mat, present.Value)
			} else {
				mat = append(mat, 0)
			}
		}
		matrix = append(matrix, mat)
	}
	print(matrix)
	return nil, nil
}

func print(matrix [][]int) {
	for _, m := range matrix {
		fmt.Println(m)
	}
}
