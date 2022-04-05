package main

import (
	"flag"
	"fmt"
)

func main() {
	v1 := flag.Bool("v", false, "usage")
	v2 := flag.Bool("vv", false, "usage")
	v3 := flag.Bool("vvv", false, "usage")
	v4 := flag.Bool("vvvv", false, "usage")
	v5 := flag.Bool("vvvvv", false, "usage")

  flag.Parse()

	vLv := getVerboseLevel(*v1, *v2, *v3, *v4, *v5)

	fmt.Printf("verbose level is %d\n", vLv)
}

func getVerboseLevel(v ...bool) int {
	cnt := 0
	for idx, i := range v {
		if i {
			cnt = idx + 1
		}
	}
	return cnt
}
