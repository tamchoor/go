package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	
	var result int

	if (idx < 0) {
		return result, errors.New("negative index")
	}
	if (len(arr) == 0) {
		return result, errors.New("empty slice")
	}
	if (len(arr) < idx) {
		return result, errors.New("index is out of bounds")
	}


	e := unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + uintptr(idx - 1) * unsafe.Sizeof(arr[0]))
	result = *(*int)(e)

	return result, nil
}

func main() {
	var err error
	var res int

	test1 := []int{3,8,9,1,4}
	fmt.Println(test1)
	i1 := 4
	res, err = getElement(test1, i1)
	fmt.Printf("i = %d, res = %d, with the error: %v\n", i1, res, err)
	i2 := -2
	res, err = getElement(test1, i2)
	fmt.Printf("i = %d, res = %d, with the error: %v\n", i2, res, err)
	i4 := 15
	res, err = getElement(test1, i4)
	fmt.Printf("i = %d, res = %d, with the error: %v\n", i4, res, err)
	test3 := []int{}
	fmt.Println(test3)
	i3 := 7
	res, err = getElement(test3, i3)
	fmt.Printf("i = %d, res = %d, with the error: %v\n", i3, res, err)
}	