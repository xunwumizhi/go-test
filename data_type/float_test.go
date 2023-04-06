package main

import (
	"fmt"
	"strconv"
	"testing"
)

// TestParseFloat 注意精度问题
func TestParseFloat(t *testing.T) {
	// str := "0.900"
	str := "0.700"
	ratio, err := strconv.ParseFloat(str, 64)
	uRatio := uint32(ratio * 100)
	fmt.Println(ratio, err, uRatio)

	ratio, err = strconv.ParseFloat(str, 32)
	uRatio = uint32(ratio * 100)
	fmt.Println(ratio, err, uRatio)
}
