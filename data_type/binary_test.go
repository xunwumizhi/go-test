package main

import (
	"fmt"
	"math"
	"testing"
)

func TestBinary(t *testing.T) {
	var _8var = 011
	var _16var = 0x11
	fmt.Println("011 to 10进制: ", _8var)
	fmt.Println("0x11 to 10进制: ", _16var)
}

func TestBitOp(t *testing.T) {
	var a uint64 = 124564
	var bitKep uint64 = 3
	fmt.Printf("%b\n", ^bitKep)
	fmt.Printf("%b & %b = %b\n", a, ^bitKep, a&(^bitKep))
	fmt.Println(a - a&(^bitKep))

}

func TestBit(t *testing.T) {
	// 取反
	var v uint32 = 474
	fmt.Printf("v:  %032b, \n^v: %032b\n", v, ^v)

	// var b uint32 = 0x1
	var b uint32 = 0x2
	fmt.Printf("^b: %b, max-b: %b\n", ^b, math.MaxUint32-b)
	fmt.Printf("max^b: %b\n", math.MaxUint32^b)

	var a uint32 = 1111
	c := a & (^b)
	fmt.Printf("a: %b\n", a)
	fmt.Printf("c: %b\n", c)
}

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
