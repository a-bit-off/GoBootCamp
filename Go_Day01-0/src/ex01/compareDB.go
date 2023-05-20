package main

import (
	"flag"
	"fmt"

	"src/ex01/Compare"
)

func main() {
	fOld := flag.String("old", "original_database.xml", "File path")
	fNew := flag.String("new", "stolen_database.json", "File path")
	flag.Parse()
	if *fOld == "" || *fNew == "" {
		fmt.Println("File path not specified")
		return
	}
	Compare.Compare("../DataBase/"+*fOld, "../DataBase/"+*fNew)
}
