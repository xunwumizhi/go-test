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
		return list[i] < list[j] // 逆序
	})
	fmt.Println(list)
}
