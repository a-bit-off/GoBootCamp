package main

import (
	"errors"
	"ex00/DBReader"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	filePath, err := ParseFileName()
	if err != nil {
		fmt.Println(err)
		return
	}

	format, err := DBReader.ChooseFormat(*filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	byt, err := DBReader.ReadFile(*filePath, format)
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
	filePath := flag.String("f", "", "File path")
	flag.Parse()

	if *filePath == "" {
		return nil, errors.New("File path not specified")
	}
	return filePath, nil
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
