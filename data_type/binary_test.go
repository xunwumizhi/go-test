package main

import (
	"fmt"
	"math"
	"testing"
)

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
