package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortSlice(t *testing.T) {
	list := []int{1, 3, 2, 6, 7}
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j] // 逆序
	})
	fmt.Println(list)

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j] // 正序
	})
	fmt.Println(list)

	sort.Ints(list)
	fmt.Println(list)
}

func TestSortFn(t *testing.T) {
	list := []int{1, 3, 2, 6, 2, 7}
	sortInt(list, func(i, j int) bool {
		// return list[i] >= list[j]
		return list[i] > list[j]
	})
	fmt.Println(list)
}

func sortInt(list []int, less func(i, j int) bool) {
	sort.Slice(list, less)
}
