package Compare

import (
	"ex01/DBReader"
	"ex01/DBReader/MyJson"
	"ex01/DBReader/MyXml"
	"fmt"
)

type recipes struct {
	Cake []cake `json:"cake"`
}

type cake struct {
	Name        string        `json:"name"`
	Time        string        `json:"time"`
	Ingredients []ingredients `json:"ingredients"`
}

type ingredients struct {
	IngredientName  string `json:"ingredient_name"`
	IngredientCount string `json:"ingredient_count"`
	IngredientUnit  string `json:"ingredient_unit,omitempty"`
}

func Compare(oldFileName, newFileName string) ([]string, error) {
	var xml, json recipes
	if x, j, err := ParseFile(oldFileName, newFileName); err != nil {
		return nil, err
	} else {
		xml = UniformXml(x)
		json = UniformJson(j)
	}
	diffFeilds := CompareRecipes(xml, json)
	return diffFeilds, nil
}

func CompareRecipes(r1, r2 recipes) []string {
	diffFeilds := make([]string, 0)
	diffIngredients := make([]string, 0)

	// мапа для хранения уникальных названий Cake
	uniqueCakeNames := make(map[string]bool)
	for _, cake1 := range r1.Cake {
		uniqueCakeNames[cake1.Name] = true
	}

	// сравниваем названия
	for _, cake2 := range r2.Cake {
		if _, ok := uniqueCakeNames[cake2.Name]; !ok {
			diffFeilds = append(diffFeilds, fmt.Sprintf("REMOVED cake \"%s\"", cake2.Name))
		} else {
			// если нашли удаляем из мапы
			delete(uniqueCakeNames, cake2.Name)

			// находим совпавшее название
			for _, cake1 := range r1.Cake {
				if cake1.Name == cake2.Name {
					// сравниваем время
					if cake1.Time != cake2.Time {
						diffFeilds = append(diffFeilds,
							fmt.Sprintf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"",
								cake2.Name, cake1.Time, cake2.Time))
					}
					// сравниваем Ingredients
					diffIngredients = append(diffIngredients, CompareIngredients(cake2, cake1)...)

				}
			}
		}
	}
	diffFeilds = append(diffFeilds, diffIngredients...)
	return diffFeilds
}

func CompareIngredients(cake1, cake2 cake) []string {
	diffIngredients := make([]string, 0)

	// мапа для хранения уникальных названий ingredients
	uniqueIngredientsName := make(map[string]bool)
	for _, ingredient1 := range cake1.Ingredients {
		uniqueIngredientsName[ingredient1.IngredientName] = true
	}

	// сравниваем названия
	for _, ingredient2 := range cake2.Ingredients {
		if _, ok := uniqueIngredientsName[ingredient2.IngredientName]; !ok {
			diffIngredients = append(diffIngredients,
				fmt.Sprintf("REMOVED ingredient \"%s\" for cake  \"%s\"",
					ingredient2.IngredientName, cake2.Name))
		} else {
			// если нашли удаляем из мапы
			delete(uniqueIngredientsName, ingredient2.IngredientName)

			// находим совпавшее название
			for _, ingredient1 := range cake1.Ingredients {
				if ingredient1.IngredientName == ingredient2.IngredientName {
					// сравниваем IngredientCount
					diffIngredients = append(diffIngredients,
						CompareIngredientCount(ingredient1, ingredient2, cake1)...)

					// сравниваем IngredientUnit
					diffIngredients = append(diffIngredients,
						CompareIngredientUnit(ingredient1, ingredient2, cake1)...)
				}
			}
		}

	}
	return diffIngredients
}

func CompareIngredientCount(ingredient1, ingredient2 ingredients, cake1 cake) []string {
	diffIngredients := make([]string, 0)
	if ingredient1.IngredientCount != "" && ingredient2.IngredientCount != "" &&
		ingredient1.IngredientCount != ingredient2.IngredientCount {
		diffIngredients = append(diffIngredients,
			fmt.Sprintf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"",
				ingredient1.IngredientName, cake1.Name, ingredient1.IngredientCount, ingredient2.IngredientCount))
	} else if ingredient1.IngredientCount != "" && ingredient2.IngredientCount == "" {
		diffIngredients = append(diffIngredients,
			fmt.Sprintf("REMOVED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\"",
				ingredient1.IngredientName, cake1.Name, ingredient1.IngredientCount))
	} else if ingredient1.IngredientCount == "" && ingredient2.IngredientCount != "" {
		diffIngredients = append(diffIngredients,
			fmt.Sprintf("ADDED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\"",
				ingredient1.IngredientName, cake1.Name, ingredient2.IngredientCount))
	}
	return diffIngredients
}

func CompareIngredientUnit(ingredient1, ingredient2 ingredients, cake1 cake) []string {
	diffIngredients := make([]string, 0)
	if ingredient1.IngredientUnit != "" && ingredient2.IngredientUnit != "" &&
		ingredient1.IngredientUnit != ingredient2.IngredientUnit {
		diffIngredients = append(diffIngredients,
			fmt.Sprintf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"",
				ingredient1.IngredientName, cake1.Name, ingredient1.IngredientUnit, ingredient2.IngredientUnit))
	} else if ingredient1.IngredientUnit != "" && ingredient2.IngredientUnit == "" {
		diffIngredients = append(diffIngredients,
			fmt.Sprintf("REMOVED unit for ingredient \"%s\" for cake \"%s\" - \"%s\"",
				ingredient1.IngredientName, cake1.Name, ingredient1.IngredientUnit))
	} else if ingredient1.IngredientUnit == "" && ingredient2.IngredientUnit != "" {
		diffIngredients = append(diffIngredients,
			fmt.Sprintf("ADDED unit for ingredient \"%s\" for cake \"%s\" - \"%s\"",
				ingredient1.IngredientName, cake1.Name, ingredient2.IngredientUnit))
	}
	return diffIngredients
}

func ParseFile(oldFileName, newFileName string) (*MyXml.Recipes, *MyJson.Recipes, error) {
	xml := MyXml.GetRecipes()

	if _, err := DBReader.ReadFile(oldFileName, xml); err != nil {
		return nil, nil, err
	}

	json := MyJson.GetRecipes()
	if _, err := DBReader.ReadFile(newFileName, json); err != nil {
		return nil, nil, err
	}

	return xml, json, nil
}

func UniformXml(xml *MyXml.Recipes) recipes {
	rec := recipes{}

	for i := 0; i < len(xml.Cake); i++ {
		Cake := cake{xml.Cake[i].Name, xml.Cake[i].Stovetime, nil}
		for j := 0; j < len(xml.Cake[i].Ingredients.Item); j++ {
			Ingredients := ingredients{xml.Cake[i].Ingredients.Item[j].Itemname,
				xml.Cake[i].Ingredients.Item[j].Itemcount,
				xml.Cake[i].Ingredients.Item[j].Itemunit}
			Cake.Ingredients = append(Cake.Ingredients, Ingredients)
		}
		rec.Cake = append(rec.Cake, Cake)
	}
	return rec
}

func UniformJson(json *MyJson.Recipes) recipes {
	rec := recipes{}

	for i := 0; i < len(json.Cake); i++ {
		Cake := cake{json.Cake[i].Name, json.Cake[i].Time, nil}
		for j := 0; j < len(json.Cake[i].Ingredients); j++ {
			Ingredients := ingredients{json.Cake[i].Ingredients[j].IngredientName,
				json.Cake[i].Ingredients[j].IngredientCount,
				json.Cake[i].Ingredients[j].IngredientUnit}
			Cake.Ingredients = append(Cake.Ingredients, Ingredients)
		}
		rec.Cake = append(rec.Cake, Cake)
	}
	return rec
}
