package main

import (
	"fmt"
	"testing"
)

// 倒转
func TestReverse(t *testing.T) {
	a := []int{0, 1, 2, 3}    // len: 4, len/2 = 2
	b := []int{0, 1, 2, 3, 4} // len: 5, len/2 = 2

	fn := func(list []int) {
		fmt.Println(len(list) / 2)
		for i := 0; i < len(list)/2; i++ { // 交换方法1
			list[i], list[len(list)-1-i] = list[len(list)-1-i], list[i]
		}
		fmt.Println(list)
	}

	fn(a)
	fmt.Println(a)

	fn(b)
	fmt.Println(b)

	fn1 := func(list []int) []int {
		for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
		return list
	}
	fmt.Println(fn1(a))

}
