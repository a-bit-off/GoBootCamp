package Compare

import (
	"fmt"
	"src/ex00/DBReader"
	"src/ex00/DBReader/MyJson"
	"src/ex00/DBReader/MyXml"

	"golang.org/x/exp/slices"
)

func Compare(fNameOld string, fNameNew string) error {
	formatOld := MyXml.GetRecipes()
	if _, err := DBReader.ReadFile(fNameOld, formatOld); err != nil {
		return err
	}

	formatNew := MyJson.GetRecipes()
	if _, err := DBReader.ReadFile(fNameNew, formatNew); err != nil {
		return err
	}

	// Матрица Cake
	cakeMatrix := make([][]bool, len(formatOld.Cake))
	for i := range cakeMatrix {
		cakeMatrix[i] = make([]bool, len(formatNew.Cake))
	}

	// заполнение матрицы Cake
	for i := 0; i < len(formatOld.Cake); i++ {
		for j := 0; j < len(formatNew.Cake); j++ {
			if formatOld.Cake[i].Name == formatNew.Cake[j].Name {
				cakeMatrix[i][j] = true
				if formatOld.Cake[i].Stovetime != formatNew.Cake[j].Time {
					fmt.Println("CHANGED cooking time for cake", formatOld.Cake[i].Name,
						formatOld.Cake[i].Stovetime, "to", formatNew.Cake[j].Time)
				}
				break
			}
		}
	}

	for i := 0; i < len(cakeMatrix); i++ {
		if !slices.Contains(cakeMatrix[i], true) {
			fmt.Println(formatOld.Cake[i].Name, "->removed")
		}
	}

	for i := 0; i < len(cakeMatrix); i++ {
		flag := true
		for j := 0; j < len(cakeMatrix[i]); j++ {
			if cakeMatrix[i][j] {
				flag = false
			}
		}
		if flag {
			fmt.Println(formatNew.Cake[i].Name, "->added")
		}
	}
	// fmt.Println(cakeMatrix)
	return nil
}

// go run compareDB.go --old original_database.xml --new stolen_database.json
