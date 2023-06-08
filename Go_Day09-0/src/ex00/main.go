package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := sleepSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	for v := range ch {
		fmt.Println(v)
	}
}

func sleepSort(numbers []int) chan int {
	lenNums := len(numbers)

	ch := make(chan int, lenNums)
	defer close(ch)

	var wg sync.WaitGroup
	wg.Add(lenNums)
	defer wg.Wait()

	for _, num := range numbers {
		go func(num int) {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(num))
			ch <- num
		}(num)
	}
	return ch
}
