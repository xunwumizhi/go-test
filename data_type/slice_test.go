package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestSliceOp(t *testing.T) {
	// delete
	a := []int{-1, 0, 5, 4}
	var b []int = make([]int, len(a))
	copy(b, a)
	fmt.Println("b", b)
	var c []int = make([]int, len(a))
	copy(c, a)
	fmt.Println("c", c)

	deleteIndex := 1
	fmt.Println("a", a)
	b = deleteWithCopy(b, deleteIndex)
	c = append(c[:deleteIndex], c[deleteIndex+1:]...)
	fmt.Println("b", b)
	fmt.Println("c", c)
}

func deleteWithCopy(a []int, i int) []int {
	copyedNum := copy(a[i:], a[i+1:])
	fmt.Println("copy num:", copyedNum)
	return a[:len(a)-1]
}

func TestList(t *testing.T) {
	list.New()
}
