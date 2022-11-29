package main

import (
	"flag"
	"fmt"
	"reflect"
)

func compare(oldR Recipes, newR Recipes) {
	if reflect.DeepEqual(oldR, newR) {
		fmt.Println("EQUAL")
		return
	}
	checkCake(oldR, newR, "ADDED")
	checkCake(newR, oldR, "REMOVED")
	checkTime(oldR, newR)
	checkIngredient(oldR, newR, "ADDED")
	checkIngredient(newR, oldR, "REMOVED")
	checkIngredientUnit(oldR, newR, "ADDED", 1)
	checkIngredientUnit(newR, oldR, "REMOVED", 2)
}

func main() {
	oldFile := flag.String("old", "", "name of .xml/.json old file")
	newFile := flag.String("new", "", "name of .xml/.json new file")
	flag.Parse()
	if len(*oldFile) != 0 && len(*newFile) != 0 {
		oldRecipes := readDB(oldFile)
		newRecipes := readDB(newFile)
		compare(oldRecipes, newRecipes)
	} else {
		fmt.Println("Usage :\n\t -new string\n\t\t name of .xml/.json new file \n \n\t -old string\n\t\t name of .xml/.json old file")
	}
}
