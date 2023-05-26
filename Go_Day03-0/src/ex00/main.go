package NOTmain

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/elastic/go-elasticsearch/v7"
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
	pathToDataSets := flag.String("dSet", "../../materials/dataTest.csv", "path to csv file")
	flag.Parse()

	// Парсим csv файл
	dataSet, err := Parse(*pathToDataSets)
	if err != nil {
		log.Fatal(err)
	}

	WorkWithElasticsearch()
}

func WorkWithElasticsearch() error {
	// создание нового клиента
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return err
	}

	// создание индекса API
	res, err := es.Indices.Create("places")
	if err != nil {
		return err
	}
	res.Body.Close()
	res
	return nil
}

func Parse(fileCSV string) ([][]string, error) {
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

func setData(data []string) Doc {
	id, _ := strconv.Atoi(data[0])
	lon, _ := strconv.ParseFloat(data[4], 64)
	lat, _ := strconv.ParseFloat(data[5], 64)
	return Doc{
		ID:       id,
		Name:     data[1],
		Address:  data[2],
		Phone:    data[3],
		Location: elastic.GeoPoint{Lon: lon, Lat: lat},
	}
}
