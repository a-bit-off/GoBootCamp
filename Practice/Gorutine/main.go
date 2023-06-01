package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go func(k int) {
			defer wg.Done()
			fmt.Printf("%d gorutine working...\n", k)
		}(i)

	}
	wg.Wait()
	fmt.Println("all done")
}
