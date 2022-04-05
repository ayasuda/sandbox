package main

import (
	"flag"
	"fmt"
)

func main() {
	var nFlag = flag.Int("n", 1234, "help message for flagname")

	flag.Parse()
	fmt.Printf("arguments num: %d\n", flag.NArg())

	fmt.Printf("arguments as []string: %v\n", flag.Args())

	fmt.Println("each arguments")
	for i, j := 0, flag.NArg(); i < j; i++ {
		fmt.Printf("\targument %d: %s\n", i, flag.Arg(i))
	}
	fmt.Printf("n sets %d\n", *nFlag)
}
