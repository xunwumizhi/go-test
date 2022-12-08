package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestEscape(t *testing.T) {
	// s := ""
	s := ``
	se := url.QueryEscape(s)
	fmt.Println(se)
	sue, _ := url.QueryUnescape(s)
	fmt.Println(sue)
	urlValue, _ := url.ParseQuery(sue)
	fmt.Println(urlValue)

	param := ""
	pue, _ := url.QueryUnescape(param)
	fmt.Println("pue: ", pue)
	urlValue, _ = url.ParseQuery(pue)
	fmt.Println(urlValue)
}
