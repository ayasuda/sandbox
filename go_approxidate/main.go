package main

import (
	"fmt"
	"time"
)

func main() {
	show("2019-10-09")
}

func show(s string) {
	p, err := approxidate(s)
	if err != nil {
		fmt.Printf("got err: %v", err)
	}
	fmt.Printf("%s parsed %s", s, p)
}

func approxidate(s string) (time.Time, error) {
	return time.Now(), nil
}
