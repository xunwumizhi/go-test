package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help bool
	user string
)

func init() {
	for i, arg := range os.Args {
		fmt.Printf("[%d]: %s\n", i, arg)
	}
	flag.BoolVar(&help, "h", false, "help")
	flag.StringVar(&user, "u", "admin", "admin name")

	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage: command [-h help]
Options:
`)
		flag.PrintDefaults()
	}
	if help {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	fmt.Printf("hello world")
}
