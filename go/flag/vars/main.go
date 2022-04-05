package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var flagBoolVar bool
	var flagIntVar int
	var flagInt64Var int64
	var flagUintVar uint
	var flagUint64Var uint64
	var flagFloat64Var float64
	var flagStringVar string
	var flagDurationVar time.Duration

	flag.BoolVar(&flagBoolVar, "bool-var-flag", false, "usage message for flag.Bool")
	flag.IntVar(&flagIntVar, "int-var-flag", 0, "usage message for flag.Int")
	flag.Int64Var(&flagInt64Var, "int64-var-flag", 0, "usage message for flag.Int64")
	flag.UintVar(&flagUintVar, "uint-var-flag", 0, "usage message for flag.Uint")
	flag.Uint64Var(&flagUint64Var, "uint64-var-flag", 0, "usage message for flag.Uint64")
	flag.Float64Var(&flagFloat64Var, "float64-var-flag", 0, "usage message for flag.Float64")
	flag.StringVar(&flagStringVar, "string-var-flag", "", "usage message for flag.String")
	flag.DurationVar(&flagDurationVar, "duration-var-flag", 0, "usage message for flag.Duration")

	flag.Parse()
	fmt.Printf("bool-var-flag has value %t\n", flagBoolVar)
	fmt.Printf("int-var-flag has value %d\n", flagIntVar)
	fmt.Printf("int64-var-flag has value %d\n", flagInt64Var)
	fmt.Printf("uint-var-flag has value %d\n", flagUintVar)
	fmt.Printf("uint64-var-flag has value %d\n", flagUint64Var)
	fmt.Printf("float64-var-flag has value %f\n", flagFloat64Var)
	fmt.Printf("string-var-flag has value '%s'\n", flagStringVar)
	fmt.Printf("duration-var-flag has value %d\n", flagDurationVar)
}
