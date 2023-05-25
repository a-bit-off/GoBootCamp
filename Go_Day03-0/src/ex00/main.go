package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/grailbio/base/tsv"
	"github.com/olivere/elastic"
)

type Doc struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Address  string           `json:"address"`
	Phone    string           `json:"phone"`
	Location elastic.GeoPoint `json:"location"`
}

type Schema struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	Name     Name     `json:"name"`
	Address  Address  `json:"address"`
	Phone    Phone    `json:"phone"`
	Location Location `json:"location"`
}

type Name struct {
	Type string `json:"type"`
}

type Address struct {
	Type string `json:"type"`
}

type Phone struct {
	Type string `json:"type"`
}

type Location struct {
	Type string `json:"type"`
}

func main() {
	pathToDataSets := flag.String("dSet", "../../materials/data.csv", "path to csv file")
	flag.Parse()
	dataSet, err := getDataSet(*pathToDataSets)
	if err != nil {
		log.Fatal(err)
	}
	for _, ds := range dataSet {
		fmt.Println(ds)
	}
}

func getDataSet(fileCSV string) ([][]string, error) {
	file, err := os.Open(fileCSV)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var byt []byte
	reader := tsv.NewReader(file)
	reader.Read(byt)

	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	// data = data[:][:]
	return data, nil
}
