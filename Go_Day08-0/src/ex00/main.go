package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func main() {
	res, err := getElement([]int{5, 6, 7, 8}, 3)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("res:", res)
}

func getElement(arr []int, idx int) (int, error) {
	if idx < 0 || idx >= len(arr) {
		return 0, errors.New("idx out of range")
	}
	const size = unsafe.Sizeof(0)
	p := unsafe.Pointer(&arr[0])
	res := *(*int)(unsafe.Pointer(uintptr(p) + size*uintptr(idx)))
	return res, nil
}
