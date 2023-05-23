package CompareFileSystem

import (
	"bufio"
	"os"
)

func CompareFS(oldFileName, newFileName string) ([]string, error) {
	oldUnique, err := Parse(oldFileName)
	if err != nil {
		return nil, err
	}

	newUnique, err := Parse(newFileName)
	if err != nil {
		return nil, err
	}

	diffFeilds := CompareMaps(oldUnique, newUnique)
	return diffFeilds, nil
}

func Parse(fileName string) (map[string]bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	uniqueFilePath := make(map[string]bool)
	for scan.Scan() {
		uniqueFilePath[scan.Text()] = true
	}

	return uniqueFilePath, nil
}

func CompareMaps(oldUnique, newUnique map[string]bool) []string {
	diffFeilds := make([]string, 0)
	for old := range oldUnique {
		flag := true
		for new := range newUnique {
			if old == new {
				flag = false
				delete(oldUnique, old)
				delete(newUnique, new)
				break
			}
		}
		if flag {
			diffFeilds = append(diffFeilds, "REMOVED "+old)
		}
	}
	for new := range newUnique {
		diffFeilds = append(diffFeilds, "ADDED "+new)
	}

	return diffFeilds
}
