package main

import "fmt"

func main() {
	s1 := make([]int, 3)
	s1[0] = 0
	s1[1] = 1
	s1[2] = 2

	s2 := s1
	s2[0] = 4
	fmt.Println(s1)
	fmt.Println(s2)

}
