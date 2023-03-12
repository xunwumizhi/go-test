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
