package min

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sort"
)

func MinCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

func removeDuplicateElement(array []int) []int {
	result := make([]int, 0, len(array))
	temp := map[int]struct{}{}
	for _, item := range array {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func checkZero(coins []int) bool {

	for _, i := range coins {
		if i <= 0 {
			return false
		}
	}
	return true
}

func countSumm(coins []int) (summ int) {
	for _, i := range coins {
		summ += i
	}
	return
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func MinCoins2(val int, coins []int) []int {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	if len(coins) == 0 || checkZero(coins) == false {
		return []int{}
	}
	sort.Ints(coins)
	removeDuplicateElement(coins)
	var resSave, resCorrect []int
	for i := 0; i < len(coins); i++ {
		resCorrect = MinCoins(val, coins[:len(coins)-i])
		if countSumm(resCorrect) == val {
			if len(resSave) == 0 {
				resSave = resCorrect
			} else if len(resSave) > len(resCorrect) {
				resSave = resCorrect
			}
		}
	}
	if countSumm(resSave) == val {
		return resSave
	}
	return []int{}
}

func MinCoins2Optimize(val int, coins []int) []int {
	if len(coins) == 0 || checkZero(coins) == false {
		return []int{}
	}
	sort.Ints(coins)
	var resSave, resCorrect []int
	resSave = []int{}
	for i := 0; i < len(coins); i++ {
		resCorrect = MinCoins(val, coins[:len(coins)-i])
		if countSumm(resCorrect) == val {
			if len(resSave) == 0 {
				resSave = resCorrect
			} else if len(resSave) > len(resCorrect) {
				resSave = resCorrect
			}
		}
	}
	return resSave
}
