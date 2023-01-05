package min

import (
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

// MinCoins2 - func accepts a necessary amount of coins.
// The output is supposed to be a slice of coins of minimal size that can be used to express the value. If some errors it returns empty slice.
//
//	-- Func sort slice of ints
//	-- In cycle call MinCoins without last elem in sclice, check results and save if this correct and have less elements
func MinCoins2(val int, coins []int) []int {
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
