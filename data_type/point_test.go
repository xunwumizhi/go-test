package main

import (
	"fmt"
	"net"
	"testing"
)

func TestNilPoint(t *testing.T) {
	var p *net.AddrError = nil
	fmt.Println(*p) // panic
}
