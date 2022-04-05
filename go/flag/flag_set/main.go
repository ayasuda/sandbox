package main

import (
	"flag"
	"fmt"
)

func main() {

	rf := flag.String("rootFlag", "default var", "usage of root flag")

	sub := flag.NewFlagSet("sub", flag.ExitOnError)
	sf := sub.String("subFlag", "default var", "usage message")

	flag.Parse()

	subCmd := flag.Arg(0)

	switch subCmd {
	case "sub":
		// flag.Arg(0) があることで、len(flag.Args()) > 0 は自明
		sub.Parse(flag.Args()[1:])
	default:
		fmt.Println("fuga")
		// noop
	}

	fmt.Printf("rootFlag is %s\n", *rf)
	fmt.Printf("root arguments as []string: %v\n", flag.Args())
	fmt.Printf("subFlag is %s\n", *sf)
	fmt.Printf("sub arguments as []string: %v\n", sub.Args())
}
