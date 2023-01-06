package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func sleepSort(array []int) <-chan int {
	var wg sync.WaitGroup
	res := make(chan int, len(array))
	for _, elem := range array {
		wg.Add(1)
		elem1 := elem
		go func() {
			time.Sleep(time.Duration(elem1) * time.Millisecond)
			res <- elem1
			wg.Done()
		}()
	}
	wg.Wait()
	close(res)
	return res
}

func test(array []int) {
	fmt.Println("array :", array)
	res := sleepSort(array)
	fmt.Printf("arra1 : ")
	for i := range res {
		fmt.Printf(" %d", i)
	}
	fmt.Printf(" \n")
	sort.Ints(array)
	fmt.Println("arra2 :", array)
}

func main() {
	var i int
	fmt.Printf("test : %d\n", i)
	i++
	test([]int{5, 2, 3, 1, 4})
	fmt.Printf("test : %d\n", i)
	i++
	test([]int{1, 1, 5, 2, 3, 1})
	fmt.Printf("test : %d\n", i)
	i++
	test([]int{1, 1, 5, 2, 5, 7, 3, 1})
}
