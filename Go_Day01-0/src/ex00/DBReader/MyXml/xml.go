package MyXml

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

type Recipes struct {
	XMLName xml.Name `xml:"recipes"`
	Cake    []Cake   `xml:"cake"`
}

type Cake struct {
	Name        string      `xml:"name"`
	Stovetime   string      `xml:"stovetime"`
	Ingredients Ingredients `xml:"ingredients"`
}

type Ingredients struct {
	Item []Item `xml:"item"`
}

type Item struct {
	Itemname  string `xml:"itemname"`
	Itemcount string `xml:"itemcount"`
	Itemunit  string `xml:"itemunit"`
}

func (r *Recipes) Parse(reader io.Reader) error {
	byteValue, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	if err := xml.Unmarshal(byteValue, r); err != nil {
		return err
	}
	return nil
}

func (r *Recipes) ConvertPP() ([]byte, error) {
	if byt, err := xml.MarshalIndent(r, "", "    "); err != nil {
		return nil, err
	} else {
		return byt, nil
	}
}
