package main

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestDecimal(t *testing.T) {
	fmt.Println("10^3: ", 1e3)
	fmt.Println("10^6: ", 1e6) // 百万
	fmt.Println("10^9: ", 1e9) // 10亿

	fmt.Println(math.Exp2(10)) // 1K
	fmt.Println(math.Exp2(20)) // 1M
	fmt.Println(math.Exp2(30)) // 1G

	// var sec int64 = math.MaxUint32 // 2106-02-07 14:28:15 +0800 CST
	var sec int64 = math.MaxInt32 // 2038-01-19 11:14:07 +0800 CST
	fmt.Println(time.Unix(sec, 0))

	fmt.Println(1000 / 8)
	fmt.Println(125 * 1e6 / math.Exp2(20))

	fmt.Println(math.Exp2(20)) // 2^20
	fmt.Println(1 << 20)
	fmt.Println(math.Log2(1048576))
}
