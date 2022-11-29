package main

import (
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	sl, d, f bool
	ext      string
}

func initFlags(flags *Flags) {
	flag.BoolVar(&flags.sl, "sl", false, "symlink display")
	flag.BoolVar(&flags.d, "d", false, "dir display")
	flag.BoolVar(&flags.f, "f", false, "file display")
	flag.StringVar(&flags.ext, "sd", "", "files with a certain extension display")
}

func main() {

}
