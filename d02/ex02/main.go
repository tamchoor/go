package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputLines := strings.Split(string(input), "\n")

	flag.Parse()
	for i := range inputLines {
		if len(flag.Args()) > 0 {
			arg := flag.Args()[1:]
			newarg := append(arg, inputLines[i])
			cmd := exec.Command(flag.Args()[0], newarg...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("%s", out)
		}
	}
}
