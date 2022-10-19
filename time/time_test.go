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
}
