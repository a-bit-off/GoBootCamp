package MyXml

import (
	"bufio"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
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

// Convert to pretty-printing
func (r *Recipes) ConvertPP() ([]byte, error) {
	if byt, err := xml.MarshalIndent(r, "", "    "); err != nil {
		return nil, err
	} else {
		return byt, nil
	}
}

func (r *Recipes) WriteToAnotherFormat(data []byte) error {
	file, err := os.Create("fromXmlTo.json")
	defer file.Close()
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
