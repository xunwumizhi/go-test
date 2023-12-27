package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestPattern(t *testing.T) {
	p1 := `[\(|（][^\)）]*[\)|）]`
	fmt.Println(p1)

	// p2 := "[\(|（][^\)）]*[\)|）]"
	// fmt.Println(p2)
}

func TestMatch(t *testing.T) {
	reg := regexp.MustCompile(`^\p{P}+$`) // 匹配英文或中文括号及其中的内容
	list := []string{
		"h(){}",
		"h(){}h",
		"(){}h",
		"(){}(",
		"(){h}(",
		"",
	}
	for _, v := range list {
		ok := reg.MatchString(v)
		fmt.Println(v, ok)
	}

}

func TestGrammar(t *testing.T) {
	greedFn := func(pattern string, text string) {
		reg := regexp.MustCompile(pattern)
		fmt.Println(reg.FindString(text))         // 匹配到的字符
		fmt.Println(reg.FindStringSubmatch(text)) // 匹配，并提取第一个
	}
	text := "---prifex:hh;suffix;suffix----"
	greedFn(`prifex:(.*);suffix`, text)
	greedFn(`prifex:(.*?);suffix`, text) // 非贪婪
}

func TestReplace(t *testing.T) {
	str := "This is (some (text)) with (parentheses) that we want to remove (1234)。这是一段（中文）括号（里面）的文字（4567）。"
	re := regexp.MustCompile(`[\(|（][^\)）]*[\)|）]`) // 匹配英文或中文括号及其中的内容
	result := re.ReplaceAllString(str, "")          // 使用空字符串替换匹配到的内容
	fmt.Println(result)                             // 输出结果：This is some  with  that we want to remove 。这是一段括号的文字。

	re = regexp.MustCompile(`[(（][^)）]*[)）]`) // 匹配英文或中文括号及其中的内容
	result = re.ReplaceAllString(str, "")     // 使用空字符串替换匹配到的内容
	fmt.Println(result)                       // 输出结果：This is some  with  that we want to remove 。这是一段括号的文字。

	// 定义需要替换的字符串
	str = "这是一个（包含括号）和😊的字符串。句号后面还有一段。"
	// 使用正则表达式匹配所有中英文括号和emoji表情
	reg := regexp.MustCompile(`[\p{P}\p{S}]`)
	str = reg.ReplaceAllString(str, "")
	// 输出替换后的字符串
	fmt.Println(str)

}
