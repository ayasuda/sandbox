package main

import (
	"flag"
	"fmt"
)

func main() {
	const (
		defaultVal = "default value"
		usage      = "usage message"
	)
	var option string
	flag.StringVar(&option, "f", defaultVal, usage)
	flag.StringVar(&option, "flagname", defaultVal, usage)

	flag.Parse()
	fmt.Println("option is", option)
}
