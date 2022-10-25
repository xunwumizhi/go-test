package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWg(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 5
	close(ch)
	v, ok := <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)

	for i := 0; i < 100; i++ {
		StopManually()
	}

	wgTest()
}

func wgTest() {

	tasks := []string{"a", "b", "c"}

	ch := make(chan string, len(tasks))

	var wg sync.WaitGroup
	for _, t := range tasks {
		wg.Add(1)

		go func(t string) {
			defer wg.Done()

			if t != "a" {
				time.Sleep(2 * time.Second)
			}

			ch <- t + " done!"
		}(t)
	}

	timeoutCh := make(chan interface{})

	go func() {
		wg.Wait()
		close(timeoutCh)
	}()

	var dataNum int = len(tasks)
	select {
	case <-timeoutCh:

	case <-time.After(time.Second):
		dataNum = len(ch)
		fmt.Println(dataNum)

	}

	var res []string
	for i := 0; i < dataNum; i++ {
		data := <-ch
		res = append(res, data)
	}

	fmt.Println(res)

}

func ConsumerProduce() {
	type pool struct {
		goods map[string]string
		sync.Mutex
	}
	mypool := &pool{}
	mypool.goods = make(map[string]string)

	var wg sync.WaitGroup
	wg.Add(2)
	go func(mypool *pool) {
		defer wg.Done()
		fmt.Printf("write: print mypool %p\n", mypool)

		mypool.Lock()
		fmt.Println("write time: ", time.Now())
		mypool.goods["name"] = "gxyu"
		mypool.Unlock()
	}(mypool)

	go func(mypool *pool) {
		defer wg.Done()
		st := time.Now()
		for {
			mypool.Lock()
			if v, ok := mypool.goods["name"]; ok {
				fmt.Printf("read: print mypool %p\n", mypool)

				fmt.Println("read time: ", time.Now())
				fmt.Println("readed name: ", v)
				mypool.Unlock()
				break
			}
			mypool.Unlock()

			if time.Since(st).Seconds() > 1 {
				fmt.Println("time out")
				break
			}
		}
	}(mypool)

	wg.Wait()
}

func StopManually() {
	stopTimer := time.NewTimer(5 * time.Second)

	ch1Stop := make(chan struct{})
	c1Ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			// case <-stopTimer.C:
			case <-ch1Stop:
				fmt.Println("ch1 stoped")
				return
			case <-c1Ticker.C:
				// fmt.Println("ch1 working")
			}
		}
	}()

	// ch2Stop := make(chan struct{})
	c2Ticker := time.NewTicker(2 * time.Second)
	go func() {
		for {
			select {
			// case <-stopTimer.C:
			// case <-ch2Stop:
			case <-ch1Stop:
				fmt.Println("ch2 stoped")
				return
			case <-c2Ticker.C:
				// fmt.Println("ch2 working")
			}
		}
	}()

	<-stopTimer.C
	close(ch1Stop)

	// time.Sleep(8 * time.Second)
	fmt.Println("main routine over")
}
