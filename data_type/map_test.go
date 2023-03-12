package main

import (
	"fmt"
	"testing"
)

func TestMapInit(t *testing.T) {
	var intMap map[int]int
	fmt.Println(intMap == nil)
	intMap = map[int]int{}
	fmt.Println(intMap == nil)
}
