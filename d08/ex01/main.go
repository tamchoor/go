package main

import (
	"fmt"
	"reflect"
	"strings"
)

type UnknownPlant struct {
    FlowerType  string
    LeafType    string
    Color       int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
    FlowerColor int
    LeafType    string
    Height      int `unit:"inches"`
}

func describePlant(data any) {
	dataValue := reflect.ValueOf(data)
	
	for i := 0; i < dataValue.NumField(); i++ {

		current := string(reflect.TypeOf(data).Field(i).Tag)
		current= strings.Replace(current, ":", "=", 1)
		current = strings.Replace(current, "\"", "", -1)
		if (current != "") {
			current = "(" + current + ")"
		}
		current = string(reflect.TypeOf(data).Field(i).Name) + current
		fmt.Printf("%s:%v\n", current, dataValue.Field(i).Interface())
		
	}

	// fmt.Println(dataValue)
	// fmt.Println(dataType)
	
}

func main() {

	tmp1 := UnknownPlant{"has a big flour", "has a lot of round leaves", 255}
	tmp2 := AnotherUnknownPlant{220, "has one big leaf", 15}
	exampleFromSubject := AnotherUnknownPlant{10, "lanceolate", 15}

// 	FlowerColor:10
// LeafType:lanceolate
// Height(unit=inches):15

	fmt.Println()
	describePlant(tmp1)
	fmt.Println()
	describePlant(tmp2)
	fmt.Println()
	describePlant(exampleFromSubject)
	fmt.Println()

}