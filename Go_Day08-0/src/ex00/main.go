/*
Здесь, в джунглях, вы можете встретить странных существ, с которыми нужно обращаться необычным способом.
Для этой задачи вам нужно написать функцию getElement(arr []int, idx int) (int, error),
которая принимает и индекс и возвращает вам элемент с этим индексом. Кажется достаточно простым, а?
Но вот одно условие - вы не можете использовать поиск по этому индексу (например arr[idx]),
разрешен только поиск по первому элементу ( arr[0]). Возможно, вам придется вспомнить немного C,
чтобы выполнить это упражнение.

В случае любого недопустимого ввода (пустой срез, отрицательный индекс, индекс за пределами)
функция должна вернуть ошибку с текстовым объяснением проблемы.
*/
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
