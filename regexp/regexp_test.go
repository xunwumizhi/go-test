package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGrammar(t *testing.T) {
	// 定义要匹配的字符串
	str := "_999999_42350985"
	// 定义正则表达式
	re := regexp.MustCompile(`_999999_(\d+)`)
	// 使用正则表达式执行匹配
	match := re.FindStringSubmatch(str)
	fmt.Printf("提取到的数字为：%v\n", match)
}
