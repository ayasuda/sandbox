package main

import (
	"flag"
	"fmt"
)

func main() {
	var flagBool = flag.Bool("bool-flag", false, "usage message for flag.Bool")
	var flagInt = flag.Int("int-flag", 0, "usage message for flag.Int")
	var flagInt64 = flag.Int64("int64-flag", 0, "usage message for flag.Int64")
	var flagUint = flag.Uint("uint-flag", 0, "usage message for flag.Uint")
	var flagUint64 = flag.Uint64("uint64-flag", 0, "usage message for flag.Uint64")
	var flagFloat64 = flag.Float64("float64-flag", 0, "usage message for flag.Float64")
	var flagString = flag.String("string-flag", "", "usage message for flag.String")
	var flagDuration = flag.Duration("duration-flag", 0, "usage message for flag.Duration")

	flag.Parse()
	fmt.Printf("bool-flag has value %t\n", *flagBool)
	fmt.Printf("int-flag has value %d\n", *flagInt)
	fmt.Printf("int64-flag has value %d\n", *flagInt64)
	fmt.Printf("uint-flag has value %d\n", *flagUint)
	fmt.Printf("uint64-flag has value %d\n", *flagUint64)
	fmt.Printf("float64-flag has value %f\n", *flagFloat64)
	fmt.Printf("string-flag has value '%s'\n", *flagString)
	fmt.Printf("duration-flag has value %d\n", *flagDuration)
}
