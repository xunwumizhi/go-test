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

	fmt.Println(strings.TrimSpace(" fa fr a "))
	fmt.Println(strings.TrimSpace("      fa fr a  "))
	fmt.Println(strings.TrimSpace("    ") == "") // true
}

func TestStringBytes(t *testing.T) {
	s := "Hello World"
	for i, n := range s {
		fmt.Println(i, n)
	}
	var b byte = s[0]
	fmt.Println(b)
}

func TestRune(t *testing.T) {
	msg := "你好\n\n   GG豪" // 特殊字符也算长度
	fmt.Println([]rune(msg))
	fmt.Println(len([]rune(msg)))
}
