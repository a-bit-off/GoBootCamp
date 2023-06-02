package main

import (
	"ex01/client/client"
	// "ex01/client/client/operations"
	"fmt"
)

func main() {
	myClient := client.NewHTTPClient(nil)
	fmt.Println("myClient:\n", myClient)
}
