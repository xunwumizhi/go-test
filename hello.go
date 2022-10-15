package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover got an error: ", err)
		}
	}()
	fmt.Println(runtime.Version())
	// panic("hello")

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("recover from goroutine,error: ", err)
			}
		}()
		panic("in goroutine")
	}()

	time.Sleep(2 * time.Second)
}
