package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

var fileName = "./text.csv"

func TestScanner(t *testing.T) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file error: %v", err)
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		name := scanner.Text()
		fmt.Println(name)
	}
}
