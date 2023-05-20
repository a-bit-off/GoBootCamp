package main

import (
	"DBReader"
	"MyJson"
	"MyXml"
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// ./DataBase/stolen_database.json
// ./DataBase/original_database.xml

func main() {
	fileName, err := ScanFileName(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	format, err := ChooseFormat(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	byt, err := ReadFile(fileName, format)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName, err = ScanFileName(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = ioutil.WriteFile(fileName, byt, 0644); err != nil {
		fmt.Println(err)
		return
	}
}

func ScanFileName(reader io.Reader) (string, error) {
	fmt.Printf("Write your file path: ")
	in := bufio.NewScanner(reader)
	in.Scan()
	fileName, err := in.Text(), in.Err()
	if err != nil {
		return fileName, err
	}
	return fileName, nil
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
