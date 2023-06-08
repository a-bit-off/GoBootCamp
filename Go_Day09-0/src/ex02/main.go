package main

import (
	"fmt"
	"sync"
)

func main() {
	ch0 := getChan(0)
	ch1 := getChan(ch0)
	ch2 := getChan(2.5)
	ch3 := getChan("tag3")
	ch4 := getChan(true)
	ch5 := getChan(0x55)
	result := multiplex(ch1, ch2, ch3, ch4, ch5)

	for res := range result {
		fmt.Println(res)
	}
}

func getChan(tag interface{}) chan interface{} {
	ch := make(chan interface{})

	go func() {
		for i := 0; i < 3; i++ {
			ch <- tag
		}
		close(ch)
	}()

	return ch
}

func multiplex(cs ...chan interface{}) chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan interface{}) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
