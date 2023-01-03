package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFmt(t *testing.T) {
	ts := time.Now().Unix()
	tt := time.Unix(ts, 0)
	fmt.Println(tt)

	// d, _ := time.ParseDuration("12h45m3s")
	d, _ := time.ParseDuration("12.5h7s")
	fmt.Println(d)
}
