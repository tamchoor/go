package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
)

const (
	M = 1
	W = 2
	L = 3
)

func initFlag() int {

	var m, w, l bool
	flag.BoolVar(&m, "m", false, "Counting characters")
	flag.BoolVar(&w, "w", false, "Counting words")
	flag.BoolVar(&l, "l", false, "Counting lines")
	flag.Parse()

	if m && !(w || l) {
		return M
	} else if !(m || w) && l {
		return L
	} else if w && !(m || l) {
		return W
	} else {
		fmt.Println("Use only one flag -m/-w/-l whith [filename]")
		os.Exit(1)

	}

	return 0
}

func countingInFile(flag int, filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	fileHandle, err := os.Open(filename)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer fileHandle.Close()
	var count int

	fileScanner := bufio.NewScanner(fileHandle)

	if flag == M {
		file, _ := os.ReadFile(filename)
		count = len(file)
		fmt.Printf("%d\t%s\n", count, filename)
	} else if flag == W {
		fileScanner.Split(bufio.ScanWords)
		for fileScanner.Scan() {
			count++
		}
		fmt.Printf("%d\t%s\n", count, filename)
	} else if flag == L {

		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			count++
		}
		fmt.Printf("%d\t%s\n", count, filename)
	}
	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	flags := initFlag()
	var wg sync.WaitGroup
	var i int
	if len(flag.Args()) != 0 {
		for i = range flag.Args() {
			wg.Add(1)
			go countingInFile(flags, flag.Args()[i], &wg)
		}
		wg.Wait()
	} else {
		fmt.Println("Use only one flag -m/-w/-l whith [filename]")
	}
}
