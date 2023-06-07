package main

import (
	"fmt"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func main() {
	describePlant(UnknownPlant{"t1", "l1", 100})
	fmt.Println()
	describePlant(AnotherUnknownPlant{200, "l1", 22})
}

func describePlant(plant any) {
	t := reflect.TypeOf(plant)
	v := reflect.ValueOf(plant)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := ""
		if string(f.Tag) != "" {
			tag += "(" + string(f.Tag) + ")"
		}
		fmt.Println(f.Name+tag+":", v.Field(i))
	}
}
