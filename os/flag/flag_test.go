package main

import (
	"flag"
	"os"
	"testing"
)

var fHelp = flag.Bool("h", false, "-h help")

func TestFlat(t *testing.T) {
	flag.Parse()
	if *fHelp {
		flag.Usage()
		os.Exit(0)
	}
}
