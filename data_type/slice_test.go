package main

import (
	"container/list"
	"fmt"
	"testing"
)

type People struct {
	Name string
	Age  int
}

func TestPointWhenSlice(t *testing.T) {
	list := []People{}
	for i := 0; i < 50; i++ {
		one := &People{Name: "Tom", Age: 3}
		list = append(list, *one)
		fmt.Println(list)
		one.Name = "Tom Gone"
		fmt.Println(list)
	}
}

func TestEdge(t *testing.T) {
	list := []int{1}
	list = list[1:]
	fmt.Println(list)
}

func TestSliceOp(t *testing.T) {
	// delete
	a := []int{-1, 0, 5, 4}
	var b []int = make([]int, len(a))
	copy(b, a)
	fmt.Println("b", b)
	var c []int = make([]int, len(a))
	copy(c, a)
	fmt.Println("c", c)

	deleteIndex := 1
	fmt.Println("a", a)
	b = deleteWithCopy(b, deleteIndex)
	c = append(c[:deleteIndex], c[deleteIndex+1:]...)
	fmt.Println("b", b)
	fmt.Println("c", c)
}

func deleteWithCopy(a []int, i int) []int {
	copyedNum := copy(a[i:], a[i+1:])
	fmt.Println("copy num:", copyedNum)
	return a[:len(a)-1]
}

func TestList(t *testing.T) {
	list.New()
}

func TestBatchSlice(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	total := len(list)
	for n := 1; n <= total; n++ {
		var batch [][]int
		fn := func(st, ed int) {
			batch = append(batch, list[st:ed])
		}
		batchSlice(total, n, fn)
		fmt.Printf("size[%d],len:%d batchSlice: %+v\n", n, len(batch), batch)
	}
}

// batchSlice fn 如果直接对原来切片保存至新切片，底层实际是同份数据
func batchSlice(total, batchSize int, fn func(st, ed int)) {
	for i := 0; i < total; {
		ed := i + batchSize
		if ed > total {
			ed = total
		}
		fn(i, ed)
		i = ed
	}
}
