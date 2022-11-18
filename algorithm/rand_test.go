package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	n := 1
	index := rand.Intn(n)
	fmt.Println(index)
}
