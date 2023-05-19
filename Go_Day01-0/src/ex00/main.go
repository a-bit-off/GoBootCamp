package main

import (
	"DBReader"
	"MyJson"
	"MyXml"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// ./DataBase/stolen_database.json
// new.json
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

	// fmt.Println(string(byt))
	err = WriteToFile(byt)
	fmt.Println("err:", err)
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

func WriteToFile(byt []byte) error {
	fileName, err := ScanFileName(os.Stdin)
	if err != nil {
		return err
	}

	// Проверка наличия файла
	_, err = os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			// Файл не существует, создаем его
			file, err := os.Create(fileName)
			if err != nil {
				return err
			}
			defer file.Close()

			writer := bufio.NewWriter(file)
			_, err = writer.Write(byt)
			if err != nil {
				return err
			}
			err = writer.Flush()
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		// Файл уже существует, открываем его для записи
		file, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		_, err = writer.Write(byt)
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}
