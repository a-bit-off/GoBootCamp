package main

import (
	"DBReader"
	"MyXml"
	"fmt"
	"os"
)

// xml
func main() {
	var x DBReader.DBReader
	x = &MyXml.Recipes{}
	xmlFile, err := os.Open("./DataBase/original_database.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlFile.Close()
	if err := x.Parse(xmlFile); err != nil {
		fmt.Println(err)
		return
	}
	if byt, err := x.ConvertPP(); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(byt))
	}
}

// json
// НЕТ ПРОВЕРИКИ НА ВАЛИДНОСТЬ JSON ФАЙЛА
// func main() {
// 	var j DBReader.DBReader
// 	j = &MyJson.StolenDB{}
// 	jsonFile, err := os.Open("./DataBase/stolen_database.json")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer jsonFile.Close()
// 	if err := j.Parse(jsonFile); err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if byt, err := j.ConvertPP(); err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		fmt.Println(string(byt))
// 	}
// }
