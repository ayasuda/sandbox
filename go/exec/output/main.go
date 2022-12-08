package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("date").Output()
	if err != nil {
		fmt.Printf("error ocurred with %v\n", err)
	}
	fmt.Printf("out ->\n%s\n", string(out))
}
