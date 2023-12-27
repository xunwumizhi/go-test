package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestRate(t *testing.T) {
	var limit rate.Limit = 1
	size := 10
	batch := 3
	r := &rateHandler{
		limiter: rate.NewLimiter(limit, size),
	}
	fmt.Println("tokens", r.limiter.Tokens())
	r.limiter.WaitN(context.TODO(), size) // 清空
	fmt.Printf("token: %f allow: %v\n", r.limiter.Tokens(), r.limiter.Allow())

	begin := time.Now()
	r.Do(batch)
	fmt.Println("do cost: ", time.Since(begin))
}

type rateHandler struct {
	limiter *rate.Limiter
}

func (r *rateHandler) Do(num int) {
	if err := r.limiter.WaitN(context.TODO(), num); err != nil {
		log.Fatalln("wait error: ", err)
	}
	fmt.Printf("do:%d things\n", num)
}
