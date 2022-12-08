package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	n := 1
	index := rand.Intn(n)
	fmt.Println(index)
}

// TestShuffle 不设置种子，每次固定
func TestShuffle(t *testing.T) {
	words := strings.Fields("ink runs from the corners of my mouth")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)

}
