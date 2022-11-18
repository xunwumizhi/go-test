package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_hashMutex(t *testing.T) {
	var mutetArr [10]*sync.Mutex
	for i := range mutetArr {
		mutetArr[i] = &sync.Mutex{}
	}

	var do = func(index, mutextIndex int) {
		m := mutetArr[mutextIndex]
		m.Lock()
		defer m.Unlock()
		fmt.Printf("index: %d, mutexIndex: %d start\n", index, mutextIndex)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("index: %d, mutexIndex: %d end\n", index, mutextIndex)
	}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			do(i, 0)
		}(i)
	}
	wg.Wait()

	fmt.Println("wait...")
	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			do(i, 2)
		}(i)
	}
	wg.Wait()
}
