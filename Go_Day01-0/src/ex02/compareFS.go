package main

import (
	"ex02/CompareFileSystem"
	"flag"
	"fmt"
)

func main() {
	fOld := flag.String("old", "snapshot1.txt", "File path")
	fNew := flag.String("new", "snapshot2.txt", "File path")
	flag.Parse()
	if *fOld == "" || *fNew == "" {
		fmt.Println("File path not specified")
		return
	}
	diffFeilds, err := CompareFileSystem.CompareFS(*fOld, *fNew)
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range diffFeilds {
		fmt.Println(s)
	}
}
