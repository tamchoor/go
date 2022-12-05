package main

import (
	"container/heap"
	"fmt"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap struct {
	array []Present
}

func (presents PresentHeap) Len() int {
	return len(presents.array)
}

func (presents PresentHeap) Less(i, j int) bool {
	if presents.array[i].Value == presents.array[j].Value {
		return presents.array[i].Size > presents.array[j].Size
	} else {
		return presents.array[i].Value < presents.array[j].Value
	}
}

func (presents PresentHeap) Swap(i, j int) {
	presents.array[i].Value, presents.array[j].Value = presents.array[j].Value, presents.array[i].Value
	presents.array[i].Size, presents.array[j].Size = presents.array[j].Size, presents.array[i].Size
}

func (presents PresentHeap) sort() {
	for i := 0; i < presents.Len(); i++ {
		for j := i; j < presents.Len(); j++ {
			if presents.Less(i, j) {
				presents.Swap(i, j)
			}
		}
	}
}

func (presents *PresentHeap) Push(x any) {
	(*presents).array = append((*presents).array, x.(Present))
}

func (presents *PresentHeap) Pop() any {
	pop := (*presents).array[len((*presents).array)-1]
	(*presents).array = (*presents).array[0 : len((*presents).array)-1]
	return pop
}

func getNCoolestPresents(ph []Present, n int) PresentHeap {
	if n < 0 || n > len(ph) {
		fmt.Println("Error : n must be n > 0 || n < len")
		return PresentHeap{nil}
	}
	var oldPh PresentHeap = PresentHeap{ph}
	var newPh PresentHeap
	oldPh.sort()

	for i, elemArr := range oldPh.array {
		if i == n {
			break
		}
		newPh.Push(elemArr)
	}
	return newPh
}

func main() {
	array := []Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}}
	presents := PresentHeap{array}
	heap.Init(&presents)

	newph := getNCoolestPresents(presents.array, 2)
	fmt.Println(newph.array)

	presents.Push(Present{6, 2})
	newph = getNCoolestPresents(presents.array, 2)
	fmt.Println(newph.array)
	newph = getNCoolestPresents(presents.array, 0)
	fmt.Println(newph.array)
	newph = getNCoolestPresents(presents.array, 12)
	fmt.Println(newph.array)
	fmt.Println(presents.array)
	presents.Pop()
	fmt.Println(presents.array)
	presents.Pop()
	fmt.Println(presents.array)
	presents.Pop()
	fmt.Println(presents.array)
	newph = getNCoolestPresents(presents.array, 1)
	fmt.Println(newph.array)

}
