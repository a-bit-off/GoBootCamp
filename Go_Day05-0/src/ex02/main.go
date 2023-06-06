package main

import (
	"container/heap"
	"errors"
	"ex02/myheap"
	"fmt"
)

func main() {
	presents := []myheap.Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}}
	if res, err := getNCoolestPresents(presents, 3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func getNCoolestPresents(presents []myheap.Present, n int) ([]myheap.Present, error) {
	desc := make([]myheap.Present, 0, n)
	if len(presents) < n || n < 0 {
		return desc, errors.New("\"n\" is not valid")
	}

	myH := &myheap.PresentHeap{}
	heap.Init(myH)
	for _, p := range presents {
		heap.Push(myH, p)
	}

	for i := 0; i < n; i++ {
		desc = append(desc, heap.Pop(myH).(myheap.Present))
	}
	return desc, nil
}
