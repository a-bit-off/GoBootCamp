package main

import (
	"DBReader"
	"MyJson"
	"MyXml"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	filePath, err := ParseFileName()
	if err != nil {
		fmt.Println(err)
		return
	}

	format, err := ChooseFormat(*filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	byt, err := ReadFile(*filePath, format)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = WriteFileAnotherFormat(*filePath, byt); err != nil {
		fmt.Println(err)
		return
	}
}

func ParseFileName() (*string, error) {
	var pathDB = "./DataBase/"
	filePath := flag.String("f", "", "File path")
	flag.Parse()

	pathDB += *filePath
	*filePath = pathDB

	if *filePath == "" {
		return nil, errors.New("File path not specified")
	}
	return filePath, nil
}

func ChooseFormat(fileName string) (DBReader.DBReader, error) {
	if strings.HasSuffix(fileName, ".json") {
		return &MyJson.StolenDB{}, nil
	} else if strings.HasSuffix(fileName, ".xml") {
		return &MyXml.Recipes{}, nil
	}
	return nil, errors.New("Unknown format")
}

func ReadFile(fileName string, format DBReader.DBReader) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err = format.Parse(file); err != nil {
		return nil, err
	}

	byt, err := format.ConvertPP()
	if err != nil {
		return nil, err
	}

	return byt, nil
}

func WriteFileAnotherFormat(filePath string, byt []byte) error {
	var outFileName string
	if strings.HasSuffix(filePath, ".json") {
		outFileName = "out.xml"
	} else if strings.HasSuffix(filePath, ".xml") {
		outFileName = "out.json"
	}

	if err := ioutil.WriteFile(outFileName, byt, 0644); err != nil {
		return err
	}
	return nil
}
