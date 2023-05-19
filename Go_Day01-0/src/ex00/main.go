package main

import (
	"DBReader"
	"MyJson"
	"fmt"
	"os"
)

func main() {
	var j DBReader.DBReader
	j = &MyJson.StolenDB{}
	jsonFile, err := os.Open("./DataBase/stolen_database.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()
	if err := j.Parse(jsonFile); err != nil {
		fmt.Println(err)
		return
	}
	if byt, err := j.ConvertPP(); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(byt))
	}
}
