package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "5")

	fmt.Println("start")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("error ocurred with %v\n", err)
	}
	fmt.Println("end")
}
