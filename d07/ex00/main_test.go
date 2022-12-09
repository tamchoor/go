package min

import (
	"reflect"
	"testing"
)

type addTest struct {
	val     int
	coins   []int
	correct []int
}

func TestMinCoins(t *testing.T) {
	var TestVars = []addTest{
		addTest{val: 6, coins: []int{1, 2, 5, 6, 9}, correct: []int{6}},
		addTest{val: 6, coins: []int{1, 2, 6, 9, 5}, correct: []int{6}},
		addTest{val: 6, coins: []int{1, 2, 6, 1, 1}, correct: []int{6}},
		addTest{val: 2, coins: []int{1, 2, 6, 1, 1}, correct: []int{2}},
	}
	for _, test := range TestVars {
		got := MinCoins(test.val, test.coins)
		want := test.correct

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

func TestMinCoins2(t *testing.T) {
	var TestVars = []addTest{
		addTest{val: 6, coins: []int{1, 2, 5, 6, 9}, correct: []int{6}},
		addTest{val: 6, coins: []int{1, 2, 6, 9, 5}, correct: []int{6}},
		addTest{val: 6, coins: []int{1, 2, 6, 1, 1}, correct: []int{6}},
		addTest{val: 2, coins: []int{1, 2, 6, 1, 1}, correct: []int{2}},
	}
	for _, test := range TestVars {
		got := MinCoins2(test.val, test.coins)
		want := test.correct

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}
