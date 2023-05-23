package DBReader

import (
	"errors"
	"io"
	"os"

	"ex00/DBReader/MyJson"
	"ex00/DBReader/MyXml"

	"strings"
)

type DBReader interface {
	Parse(io.Reader) error
	ConvertPP() ([]byte, error)
}

func ChooseFormat(fileName string) (DBReader, error) {
	if strings.HasSuffix(fileName, ".json") {
		return &MyJson.Recipes{}, nil
	} else if strings.HasSuffix(fileName, ".xml") {
		return &MyXml.Recipes{}, nil
	}
	return nil, errors.New("Unknown format")
}

func ReadFile(fileName string, format DBReader) ([]byte, error) {
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
