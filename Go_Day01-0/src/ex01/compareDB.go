package main

import (
	"ex01/Compare"
	"flag"
	"fmt"
)

func main() {
	fOld := flag.String("old", "original_database.xml", "File path")
	fNew := flag.String("new", "stolen_database.json", "File path")
	flag.Parse()
	if *fOld == "" || *fNew == "" {
		fmt.Println("File path not specified")
		return
	}
	diffFeilds, err := Compare.Compare("./DataBase/"+*fOld, "./DataBase/"+*fNew)
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range diffFeilds {
		fmt.Println(s)
	}
}
