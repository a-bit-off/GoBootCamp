package main

import (
	"fmt"
	"os"
	p "parser"
)

func main() {

	// read from file
	if f, err := os.Open("./test/test.txt"); err == nil {
		if num, err := p.Parser(f); err == nil {
			fmt.Println(num)
			fmt.Println(err)
		}
	}

	// read stdin
	// if num, err := p.Parser(os.Stdin); err == nil {
	// 	fmt.Println(num)
	// 	fmt.Println(err)
	// }

}
