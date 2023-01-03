package main

import (
	"fmt"
	"math"
	"testing"
)

func TestBitCal(t *testing.T) {
	var u int64 = 1
	fmt.Println(u << 2)
	fmt.Println("1<<63", u<<63)
	fmt.Println("-1<<63", (-1)<<63)
	fmt.Println("1<<63-1", u<<63-1) // 最大正数
	fmt.Println(math.MaxInt64)
	fmt.Println("-(1<<63-1)", -(u<<63 - 1))   // 最大正数取反
	fmt.Println("-(1<<63-1)-1", -(u<<63-1)-1) // 最大正数取反后-1 为最小负数
	fmt.Println(math.MinInt64)
}
