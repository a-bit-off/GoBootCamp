package main

import (
	"ex00/myES"
	"flag"
	"log"
)

func main() {
	dataSetPath := flag.String("dataSetPath", "../../materials/data.csv", "path to csv file")
	indexName := flag.String("indexName", "places", "index name")
	dataCount := flag.Int("dataCount", 0, "number of data entries to upload")

	flag.Parse()

	if err := myES.MyES(*dataSetPath, *indexName, *dataCount); err != nil {
		log.Fatal(err)
	}
}
