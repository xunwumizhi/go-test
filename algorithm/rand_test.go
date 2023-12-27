package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	n := 1
	index := rand.Intn(n)
	fmt.Println(index)
}

// TestShuffle 不设置种子，每次固定
func TestShuffle(t *testing.T) {
	words := strings.Fields("ink runs from the corners of my mouth")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)

}

func TestRandPick(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(RandIndexs(1000000, 5000))
	fmt.Println(RandIndexs(10, 5))
	fmt.Println(RandIndexs(5, 5))
	fmt.Println(RandIndexs(5, 10))
}

// RandIndexs 随机选取下标
func RandIndexs(total int, size int) (indexMap map[int]struct{}) {
	indexMap = make(map[int]struct{}) // 下标去重
	defer func() {
		var indexList []int
		for k := range indexMap {
			indexList = append(indexList, k)
		}
		fmt.Println(indexList)
	}()

	if size >= total { // 不大于期望
		for i := 0; i < total; i++ {
			_, ok := indexMap[i]
			if ok {
				log.Println("hit", i)
			}
			indexMap[i] = struct{}{} // 返回所有
		}
		return
	}

	for len(indexMap) < size {
		index := rand.Intn(total) // (0, total]
		indexMap[index] = struct{}{}
	}
	return
}
