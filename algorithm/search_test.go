package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {
	target := 3
	// list := []int{1, 2, 4, 4, 5, 5, 6}
	list := []int{}
	index := sort.Search(len(list), func(i int) bool {
		return list[i] >= target
	})
	fmt.Println(index)
	if index == len(list) { // 没找到
		list = append(list, target)
	} else if list[index] != target { // 找到了满足条件的最小值，target可以直接插入此下标保持列表有序
		list = append(list, 0)
		copy(list[index+1:], list[index:]) // 指定位置插入元素
		list[index] = target
	}
	fmt.Println(list)
}
