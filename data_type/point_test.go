package main

import (
	"fmt"
	"net"
	"testing"
)

func TestNilPoint(t *testing.T) {
	var p *net.AddrError = nil
	fmt.Printf("print: %+v\n", p)
	// fmt.Println(*p) // panic
}

func TestDoNilStruct(t *testing.T) {
	f := &foo{}
	f.doNil()
}

type foo struct {
	err  error
	serr *net.AddrError
}

func (f *foo) doNil() (e *net.AddrError, code int) {
	defer func() {
		fmt.Printf("f.err == nil: %v, f.err: %v\n", f.err == nil, f.err)
		fmt.Printf("f.serr == nil: %v, f.serr: %v\n", f.serr == nil, f.serr)
		f.err = e
		f.serr = e
		fmt.Printf("f.err == nil: %v, f.err: %v\n", f.err == nil, f.err)     // false, error interface had type
		fmt.Printf("f.serr == nil: %v, f.serr: %v\n", f.serr == nil, f.serr) // true, struct nil point
	}()

	return nil, 0
}
