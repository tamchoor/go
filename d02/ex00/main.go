package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Flags struct {
	sl, d, f, ex bool
	dir, ext     string
}

var flags Flags

func initFlags(flags *Flags) {
	flag.BoolVar(&flags.sl, "sl", false, "symlink display")
	flag.BoolVar(&flags.d, "d", false, "dir display")
	flag.BoolVar(&flags.f, "f", false, "file display")
	flag.BoolVar(&flags.ex, "ext", false, "file display")
	flag.Parse()
	if flags.ex == true && flags.f == false {
		fmt.Println("Use flag -ext [extension] whith flag -f")
		os.Exit(1)
	} else if flags.ex == true {
		if len(flag.Args()) == 2 {
			flags.ext = flag.Args()[0]
			flags.dir = flag.Args()[1]
		} else {
			fmt.Println("Use ./program -sl -d -f -ext [extension] [dir]")
			os.Exit(1)
		}
	} else {
		if len(flag.Args()) == 1 {
			flags.dir = flag.Args()[0]
		} else {
			fmt.Println("Use ./program -sl -d -f -ext [extension] [dir]")
			os.Exit(1)
		}
	}

	if flags.sl == false && flags.d == false && flags.f == false {
		flags.sl = true
		flags.d = true
		flags.f = true
	}
}

func checkPath(flags *Flags) {
	inputDir, err := os.Open(flags.dir)
	if flags.dir == "" || err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = inputDir.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isHiddenFile(filename string) bool {
	if filename[0] == '/' {
		return filename[1] == '.'
	}
	return filename[0] == '.'
}

func myWalkFunc(path string, info os.FileInfo, err error) error {
	if err == nil {
		subPath := strings.TrimPrefix(path, flags.dir)
		if subPath != "" && !isHiddenFile(subPath) {
			if info.Mode()&(1<<2) != 0 {
				if flags.sl && info.Mode().Type() == os.ModeSymlink {
					realPath, err := filepath.EvalSymlinks(path)
					if err != nil {
						fmt.Println(path, "-> [broken]")
					} else {
						fmt.Println(path, "->", realPath)
					}
				} else if flags.d && info.IsDir() {
					fmt.Println(path)
				} else if flags.f && info.Mode().IsRegular() {
					if flags.ex == false || flags.ex == true && filepath.Ext(path) == "."+flags.ext {
						fmt.Println(path)
					}
				}

			}
		}
	}
	return nil
}

func main() {
	initFlags(&flags)
	checkPath(&flags)
	filepath.Walk(flags.dir, myWalkFunc)
}
