package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	s := ` hal 
  fasd		fasr
arf
  fasr   				
`
	list := strings.Fields(s)
	for i, v := range list {
		fmt.Printf("%d: %s, len: %d\n", i, v, len(v))
	}
}

func TestStringBytes(t *testing.T) {
	s := "Hello World"
	for i, n := range s {
		fmt.Println(i, n)
	}
	var b byte = s[0]
	fmt.Println(b)
}
