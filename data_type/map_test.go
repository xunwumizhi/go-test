package main

import (
	"fmt"
	"testing"
)

func TestMapInit(t *testing.T) {
	var intMap map[int]int
	fmt.Println(intMap == nil)
	delete(intMap, 1)

	// empty map
	v, ok := intMap[1]
	fmt.Println(v, ok)
	// intMap[1] = 3 // panic

	intMap = map[int]int{}
	fmt.Println(intMap == nil)
}
