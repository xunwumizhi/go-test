package main

import (
	"fmt"
	"reflect"
)

func main() {
	equal()
}

func equal() {
	m1 := map[string]string{
		"name": "a",
		"like": "cat",
	}
	m2 := map[string]string{
		"name": "a",
		"like": "dog",
	}
	m3 := m1
	m4 := map[string]string{
		"like": "cat",
		"name": "a",
	}

	fmt.Println(reflect.DeepEqual(m1, m2))
	fmt.Println(reflect.DeepEqual(m1, m3))
	fmt.Println(reflect.DeepEqual(m1, m4))
}
