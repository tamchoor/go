package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func readDbtoMap(oldFile *string) map[string]int {
	file, err := os.Open(*oldFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	oldDb := make(map[string]int)
	for scanner.Scan() {
		oldDb[scanner.Text()] = 1
	}
	return oldDb
}

func comparePerLine(oldDb map[string]int, newFile *string) {
	file, err := os.Open(*newFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if oldDb[scanner.Text()] == 0 {
			fmt.Println("ADDED ", scanner.Text())
		} else if oldDb[scanner.Text()] == 1 {
			oldDb[scanner.Text()] = 2
		}
	}
	for key, v := range oldDb {
		if v == 1 {
			fmt.Println("REMOVED ", key)
		}
	}
}

func main() {
	oldFile := flag.String("old", "", "name of old file")
	newFile := flag.String("new", "", "name of new file")
	flag.Parse()
	if len(*oldFile) != 0 && len(*newFile) != 0 {
		oldDb := readDbtoMap(oldFile)
		comparePerLine(oldDb, newFile)
	} else {
		fmt.Println("Usage :\n\t -new string\n\t\t name of new file \n \n\t -old string\n\t\t name of old file")
	}
}
