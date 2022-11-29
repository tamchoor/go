package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path"
)

type Recipes struct {
	Cake []Cake `json:"cake" xml:"cake"`
}

type Cake struct {
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
}

type Ingredients struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit"`
}

func printRecipes(recipes Recipes) {
	for i := 0; i < len(recipes.Cake); i++ {
		fmt.Println("name", recipes.Cake[i].Name)
		fmt.Println("time", recipes.Cake[i].Time)
		for j := 0; j < len(recipes.Cake[i].Ingredients); j++ {
			fmt.Println("\tingredient_name", recipes.Cake[i].Ingredients[j].Name)
			fmt.Println("\tingredient_count", recipes.Cake[i].Ingredients[j].Count)
			fmt.Println("\tingredient_unit", recipes.Cake[i].Ingredients[j].Unit)
		}
	}
}

type Dbreader interface {
	read() (Recipes, error)
	convert(recipes Recipes) (string, error)
}

func checkFormat(name *string) int {
	if path.Ext(*name) == ".json" {
		return 1
	} else if path.Ext(*name) == ".xml" {
		return 2
	} else {
		return 0
	}
}

func readDB(file *string) Recipes {
	format := checkFormat(file)
	var dbreader Dbreader
	if format == 1 {
		dbreader = Json(*file)
	} else if format == 2 {
		dbreader = Xml(*file)
	} else {
		fmt.Println("Error not .xml/.json file ", file)
		os.Exit(1)
	}
	recipes, err := dbreader.read()
	if err != nil {
		os.Exit(1)
	}
	return recipes
}

type Xml string

func (name Xml) read() (Recipes, error) {
	file, err := os.ReadFile(string(name))
	if err != nil {
		fmt.Println(err)
		return Recipes{}, err
	}
	var recipes Recipes
	err = xml.Unmarshal(file, &recipes)
	if err != nil {
		fmt.Println("Error xml.Unmarshal :", name, err)
	}
	return recipes, err
}

func (name Xml) convert(recipes Recipes) (string, error) {
	convertByte, err := json.MarshalIndent(recipes, "", "    ")
	if err != nil {
		fmt.Println("Error xml.MarshalIndent:", name, err)
		return "", err
	}
	return string(convertByte), err
}

type Json string

func (name Json) read() (Recipes, error) {
	file, err := os.ReadFile(string(name))
	if err != nil {
		fmt.Println(err)
		return Recipes{}, err
	}
	var recipes Recipes
	err = json.Unmarshal(file, &recipes)
	if err != nil {
		fmt.Println("Error xml.Unmarshal :", name, err)
	}
	return recipes, err
}

func (name Json) convert(recipes Recipes) (string, error) {
	convertByte, err := xml.MarshalIndent(recipes, "", "    ")
	if err != nil {
		fmt.Println("Error xml.MarshalIndent:", name, err)
		return "", err
	}
	return string(convertByte), err
}

func findCake(recipes Recipes, name string) int {
	for i, cake := range recipes.Cake {
		if cake.Name == name {
			return i
		}
	}
	return -1
}

func findIngredient(cake Cake, ingredient string) int {

	for i, s := range cake.Ingredients {
		if ingredient == s.Name {
			return i
		}
	}
	return -1
}

func checkCake(oldR Recipes, newR Recipes, str string) {

	for _, cake := range newR.Cake {
		if findCake(oldR, cake.Name) == -1 {
			fmt.Printf(str)
			fmt.Printf(" cake \"%s\"\n", cake.Name)
		}
	}
}

func checkTime(oldR Recipes, newR Recipes) {
	for _, cake := range newR.Cake {
		indx := findCake(oldR, cake.Name)
		if indx != -1 {
			if oldR.Cake[indx].Time != cake.Time {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", cake.Name, cake.Time, oldR.Cake[indx].Time)
			}
		}
	}
}

func checkIngredient(oldR Recipes, newR Recipes, str string) {
	for _, cake := range newR.Cake {
		indx := findCake(oldR, cake.Name)
		if indx != -1 {
			for _, ingrdnt := range cake.Ingredients {
				igrdntIndx := findIngredient(oldR.Cake[indx], ingrdnt.Name)
				if igrdntIndx == -1 {
					fmt.Printf(str)
					fmt.Printf(" ingredient \"%s\" for cake  \"%s\"\n", ingrdnt.Name, oldR.Cake[indx].Name)
				}
			}
		}
	}
}

func checkIngredientUnit(oldR Recipes, newR Recipes, str string, flag int) {
	for _, cake := range newR.Cake {
		indx := findCake(oldR, cake.Name)
		if indx != -1 {
			for _, ingrdnt := range cake.Ingredients {
				igrdntIndx := findIngredient(oldR.Cake[indx], ingrdnt.Name)
				if igrdntIndx != -1 {

					if len(ingrdnt.Count) != 0 {
						if ingrdnt.Count != oldR.Cake[indx].Ingredients[igrdntIndx].Count && flag == 1 {
							fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" -  \"%s\" instead of \"%s\" \n", ingrdnt.Name, oldR.Cake[indx].Name, ingrdnt.Count, oldR.Cake[indx].Ingredients[igrdntIndx].Count)
						} else if len(oldR.Cake[indx].Ingredients[igrdntIndx].Count) == 0 {
							fmt.Printf(str)
							fmt.Printf(" unit count for ingredient \"%s\" for cake  \"%s\" -  \"%s\" \n", ingrdnt.Name, oldR.Cake[indx].Name, ingrdnt.Count)
						}
					}

					if len(ingrdnt.Unit) != 0 {
						if ingrdnt.Unit != oldR.Cake[indx].Ingredients[igrdntIndx].Unit && flag == 1 {
							fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" -  \"%s\" instead of \"%s\" \n", ingrdnt.Name, oldR.Cake[indx].Name, ingrdnt.Unit, oldR.Cake[indx].Ingredients[igrdntIndx].Unit)
						} else if len(oldR.Cake[indx].Ingredients[igrdntIndx].Unit) == 0 {
							fmt.Printf(str)
							fmt.Printf(" unit \"%s\" for ingredient \"%s\" for cake  \"%s\" \n", ingrdnt.Unit, ingrdnt.Name, oldR.Cake[indx].Name)
						}
					}
				}
			}
		}
	}

}
