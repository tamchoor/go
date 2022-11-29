package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

func checkFormat(name *string) int {
	if path.Ext(*name) == ".json" {
		return 1
	} else if path.Ext(*name) == ".xml" {
		return 2
	} else {
		return 0
	}
}

func startConvert(dbreader Dbreader) {
	recipes, err := dbreader.read()
	if err != nil {
		os.Exit(1)
	}
	convertString, err := dbreader.convert(recipes)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(convertString)
}

func main() {
	file := flag.String("f", "", "name of .xml/.json file")
	flag.Parse()
	format := checkFormat(file)
	if format != 0 {
		if format == 1 {
			startConvert(Json(*file))
		} else if format == 2 {
			startConvert(Xml(*file))
		}
	} else {
		fmt.Println("Usage :\n\t-f string\n\t\tname of .xml/.json file")
	}
}
