/*
create new post
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	Header  string `json:"header"`
	Content string `json:"content"`
}

var header = "Мозазавры"
var content = `
	Мозазавры (лат. Mosasauridae) — семейство вымерших морских ящеров (Lacertilia)
из надсемейства мозазавроидей (Mosasauroidea или Mosasauria). Довольно близкие
родственники современных варанов.

Первые ископаемые остатки мозазавров были найдены в известняковом карьере в
Маастрихте на Маасе в 1764 году. В настоящее время ископаемые остатки представителей
этого семейства уже были обнаружены на всех континентах, включая Антарктиду.

Мозазавры представляли собой большую группу верхнемеловых морских рептилий,
главным образом хищников крупных или средних размеров. Большинство известных
видов населяло тёплые, мелководные моря, широко распространённые в позднем
меловом периоде. В силу своей специализации, мозазавры очень сильно отличались
от современных ящериц: строение их внутренних органов больше напоминало таковое
у китообразных, они были живородящими, имели высокие темпы обмена веществ
и были теплокровными.

Мозазавры, вероятно, произошли от вымершего семейства водных ящериц айгиалозаврид
(Aigialosauridae). В течение последних 20 млн лет мелового периода
(туронский — маастрихтский века) мозазавры вытеснили всех своих конкурентов
в лице крупных ламнообразных акул и последних плиозавров, и стали доминирующими
морскими хищниками своего времени. Они исчезли вместе с динозаврами и птерозаврами
в результате массового вымирания в конце мелового периода, произошедшего 66 млн лет назад.`

var baseURL = "http://localhost:8888/post/new-post"

func main() {
	// create new Post
	post := Post{header, content}

	// encode a1 to bytes
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(post)

	// create new client
	client := http.Client{}

	// send request with new post data and get response
	resp, err := client.Post(baseURL, "application/json", bytes.NewBuffer(data.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read response
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print response
	fmt.Println(string(rb))
}
