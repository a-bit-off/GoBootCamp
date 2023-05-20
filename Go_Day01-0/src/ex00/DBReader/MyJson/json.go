package MyJson

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type StolenDB struct {
	Cake []Cake `json:"cake"`
}

type Cake struct {
	Name        string        `json:"name"`
	Time        string        `json:"time"`
	Ingredients []Ingredients `json:"ingredients"`
}

type Ingredients struct {
	IngredientName  string `json:"ingredient_name"`
	IngredientCount string `json:"ingredient_count"`
	IngredientUnit  string `json:"ingredient_unit,omitempty"`
}

func (s *StolenDB) Parse(reader io.Reader) error {
	byteValue, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(byteValue, s); err != nil {
		return err
	}
	return nil
}

// Convert to pretty-printing
func (s *StolenDB) ConvertPP() ([]byte, error) {
	if byt, err := json.MarshalIndent(s, "", "    "); err != nil {
		return nil, err
	} else {
		return byt, nil
	}
}
