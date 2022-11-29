package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
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
