package min

import (
	// "fmt"

	"reflect"
	"testing"
)

type addTest struct {
	val     int
	coins   []int
	correct []int
}

var TestVars = []addTest{
	addTest{val: 6, coins: []int{1, 2, 5, 6, 9}, correct: []int{6}},
	addTest{val: 6, coins: []int{1, 2, 6, 9, 5}, correct: []int{6}},
	addTest{val: 6, coins: []int{1, 2, 6, 1, 1}, correct: []int{6}},
	addTest{val: 2, coins: []int{1, 2, 6, 1, 1}, correct: []int{2}},
	addTest{val: 25, coins: []int{5, 10, 24}, correct: []int{10, 10, 5}},
	addTest{val: 13, coins: []int{1, 5, 10}, correct: []int{10, 1, 1, 1}},
	addTest{val: 13, coins: []int{1, -5, 10}, correct: []int{}},
	addTest{val: 13, coins: []int{0, 5, 10}, correct: []int{}},
	addTest{val: 206, coins: []int{1, 2, 6, 1, 1, 10, 16, 14, 3, 8, 8}, correct: []int{16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 14}},
}

func TestMinCoins2(t *testing.T) {
	for _, test := range TestVars {
		got := MinCoins2(test.val, test.coins)
		want := test.correct

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

func TestMinCoins2Optimize(t *testing.T) {
	for _, test := range TestVars {
		got := MinCoins2Optimize(test.val, test.coins)
		want := test.correct

		if reflect.DeepEqual(got, want) == false {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

// var j int

// func BenchmarkMinCoins2(b *testing.B) {
// 	if j < 8 {
// 		fmt.Println(TestVars[j])
// 		for i := 0; i < b.N; i++ {
// 			MinCoins2(TestVars[j].val, TestVars[j].coins)
// 		}
// 		j++
// 	}
// }

func BenchmarkMinCoins20(b *testing.B) {
	j := 0
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins21(b *testing.B) {
	j := 1
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins22(b *testing.B) {
	j := 2
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins23(b *testing.B) {
	j := 3
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins24(b *testing.B) {
	j := 4
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins25(b *testing.B) {
	j := 5
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins26(b *testing.B) {
	j := 6
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins27(b *testing.B) {
	j := 7
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins20Optimize(b *testing.B) {
	j := 0
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2Optimize(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins21Optimize(b *testing.B) {
	j := 1
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2Optimize(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins22Optimize(b *testing.B) {
	j := 2
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2Optimize(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins23Optimize(b *testing.B) {
	j := 3
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2Optimize(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins24Optimize(b *testing.B) {
	j := 4
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2Optimize(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins25Optimize(b *testing.B) {
	j := 5
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2Optimize(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins26Optimize(b *testing.B) {
	j := 6
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2Optimize(TestVars[j].val, TestVars[j].coins)
	}
}

func BenchmarkMinCoins27Optimize(b *testing.B) {
	j := 7
	// fmt.Println(TestVars[j])
	for i := 0; i < b.N; i++ {
		MinCoins2Optimize(TestVars[j].val, TestVars[j].coins)
	}
}
