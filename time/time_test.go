package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeCal(t *testing.T) {
	nowT := time.Now()
	fmt.Println(nowT)

	newNowT := nowT.Add(24 * time.Hour)
	fmt.Println(nowT) // 不变
	fmt.Println(newNowT)
}

func TestTimeFmt(t *testing.T) {
	ts := time.Now().Unix()
	tt := time.Unix(ts, 0)
	fmt.Println(tt)

	// d, _ := time.ParseDuration("12h45m3s")
	d, _ := time.ParseDuration("12.5h7s")
	fmt.Println(d)
}

func TestParse(t *testing.T) {
	now := time.Now().Format(time.RFC3339)
	fmt.Println(now)

	fmt.Println(time.Parse(time.RFC3339, "2023-10-08 12:00:01")) // error
	fmt.Println(time.Parse(time.RFC3339, "2023-10-08T12:00:01+08:00"))
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	var cnt int
	for {
		t := <-ticker.C
		fmt.Println("ticker: ", t)

		cnt++
		if cnt >= 3 {
			break
		}
	}

	cnt = 0
	for t := range ticker.C {
		fmt.Println("ticker in for range: ", t)
		cnt++
		if cnt >= 3 {
			break
		}
	}

	cnt = 0
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("ticker in select: ", t)
			cnt++
		}
		if cnt >= 3 {
			break
		}
	}
}
