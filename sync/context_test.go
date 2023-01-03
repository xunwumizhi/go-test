package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TestCancelCtx context.WithCancel
func TestCancelCtx(t *testing.T) {
	p := context.TODO()
	pCtx, pCancel := context.WithCancel(p)
	ctx, cancel := context.WithCancel(pCtx)
	go func(context.Context) {
		time.Sleep(time.Second)
		cancel()
	}(ctx)

	<-ctx.Done()
	fmt.Println("child done: ", ctx.Err())
	select {
	case <-pCtx.Done():
		fmt.Println("parent done: ", pCtx.Err())
	default:
		fmt.Println("parent not done") // 不影响父节点
	}

	pCancel() // 主动关闭父节点
	fmt.Println("cancel parent")
	select {
	case <-pCtx.Done():
		fmt.Println("parent done: ", pCtx.Err())
	default:
		fmt.Println("parent not done")
	}
}

func TestTimeoutCtx(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	go func(context.Context) {
		time.Sleep(3 * time.Second) // context deadline exceeded
		// time.Sleep(1 * time.Second) // canceled
		cancel()
	}(ctx)

	<-ctx.Done()
	fmt.Println("ctx error: ", ctx.Err())
}
