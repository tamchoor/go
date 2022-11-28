package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"sort"
)

type Inputs struct {
	n   int
	err error
}

type Flags struct {
	mean, median, mode, sd bool
}

type ArrayInfo struct {
	mean, median, sd float64
	mode             int
}

func readFromInputs() []int {
	var array []int
	var inpt Inputs
	var tmp int
	for inpt.n, inpt.err = fmt.Scanln(&tmp); inpt.err != io.EOF; inpt.n, inpt.err = fmt.Scanln(&tmp) {
		if inpt.err == nil {
			if tmp < -100000 || tmp > 100000 {
				print("Error: it is not between -100000 and 100000\n")
			} else {
				array = append(array, tmp)
			}
		} else {
			print("Error: it is not int\n")
		}
	}
	return array
}

func findMedian(array []int) float64 {
	var len int = len(array)
	var median float64
	if len%2 == 0 {
		median = (float64(array[len/2]) + float64(array[(len/2)-1])) / 2
	} else {
		median = float64(array[len/2])
	}
	return median
}

func findMode(array []int) int {
	var mode, counter, tmpmode, tmpcounter int
	mode = array[0]
	tmpmode = array[0]
	for _, elem := range array {
		if tmpmode == elem {
			tmpcounter++
		} else if tmpmode != elem {
			tmpcounter = 1
			tmpmode = elem
		}
		if tmpcounter > counter {
			counter = tmpcounter
			mode = tmpmode
		}
	}
	return mode
}

func findMean(array []int) float64 {
	var mean float64
	var sum int
	for _, elem := range array {
		sum += elem
	}
	mean = float64(sum) / float64(len(array))
	return mean
}

func findSD(array []int, mean float64) float64 {
	var sd float64
	for _, elem := range array {
		sd += math.Pow(float64(elem)-mean, 2)
	}
	sd = math.Sqrt(sd / float64(len(array)))
	return sd
}

func initFlags(flags *Flags) {
	flag.BoolVar(&flags.mean, "mean", false, "not display")
	flag.BoolVar(&flags.median, "median", false, "not display")
	flag.BoolVar(&flags.mode, "mode", false, "not display")
	flag.BoolVar(&flags.sd, "sd", false, "not display")
}

func printInfo(array []int, flags Flags) {
	sort.Ints(array)
	var arInfo ArrayInfo
	arInfo.mean = findMean(array)

	if flags.mean {
		fmt.Printf("Mean: %.2f\n", arInfo.mean)
	}
	if flags.median {
		arInfo.median = findMedian(array)
		fmt.Printf("Median: %.2f\n", arInfo.median)
	}
	if flags.mode {
		arInfo.mode = findMode(array)
		fmt.Printf("Mode: %d\n", arInfo.mode)
	}
	if flags.sd {
		arInfo.sd = findSD(array, arInfo.mean)
		fmt.Printf("SD: %.2f\n", arInfo.sd)
	}
}

func main() {
	var flags Flags
	initFlags(&flags)
	flag.Parse()
	if !flags.mean && !flags.median && !flags.mode && !flags.sd {
		flags.mean = true
		flags.median = true
		flags.mode = true
		flags.sd = true
	}

	var array []int = readFromInputs()
	if len(array) == 0 {
		print("array is empty\n")
		return
	}
	printInfo(array, flags)
}
