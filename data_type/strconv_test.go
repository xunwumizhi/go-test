package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestZero(t *testing.T) {
	str := ""
	i, err := strconv.Atoi(str)
	fmt.Printf("i: %d, err: %v\n", i, err)
}
