package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "this is custom message\n")
	}
	flag.Parse()
}
