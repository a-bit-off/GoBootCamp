package Compare

import (
	"fmt"
	"src/ex00/DBReader"
	"src/ex00/DBReader/MyJson"
)

func Compare(fNameOld string, fNameNew string) error {
	formatOld := &MyJson.StolenDB{}
	// formatNew, err := DBReader.ChooseFormat(newfName)
	// if err != nil {
	// 	return err
	// }
	if _, err := DBReader.ReadFile(fNameOld, formatOld); err != nil {
		return err
	}
	fmt.Println(formatOld)

	// fmt.Println(formatOld.GetStruct())
	// bytNew, err := DBReader.ReadFile(newfName, formatNew)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(string(bytOld), "\n\n\n\n")
	// fmt.Println(string(bytNew))
	return nil
}

// go run compareDB.go --old original_database.xml --new stolen_database.json
