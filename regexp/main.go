package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	reg := regexp.MustCompile(`^\s+$`)
	test := []string{
		" fasd fa",
		"    ",
		"  \t \r \n ",
	}
	for _, t := range test {
		fmt.Println(t, ": ", reg.MatchString(t))
	}

	// s := "dw_yuguoxiu"
	// s = s[3:]
	// fmt.Println(s)
	// test0()
	// regTest()
}

func test0() {

	testStr := "1 OR 1"

	exp := `.*( OR | or |=).*`
	matched, err := regexp.MatchString(exp, testStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("matched result: ", matched)
}

func regTest() {
	target := "New size: 15; reason: All metrics below target"
	exp := `New size: .*;`
	reg := regexp.MustCompile(exp)
	want := reg.FindString(target)
	want = strings.ReplaceAll(want, "New size: ", "")
	want = strings.ReplaceAll(want, ";", "")
	fmt.Println(want)

	reg = regexp.MustCompile(`New size: (.*);`)
	gets := reg.FindStringSubmatch(target)
	if gets != nil {
		fmt.Println("submatch: 0: ", gets[0])
		fmt.Println("submatch: 1: ", gets[1])
	}

	// panic, 不是标准的语法
	// reg = regexp.MustCompile(`(?<=New size: ).*?(?=;)`)

}
