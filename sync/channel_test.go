package main

import (
	"fmt"
	"testing"
)

func TestBufferChan(t *testing.T) {
	errCh := make(chan int, 10)
	fmt.Println(len(errCh))
	errCh <- 1
	fmt.Println(len(errCh))
	close(errCh)
	fmt.Println(len(errCh))
}
