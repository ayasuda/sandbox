package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("sleep", "5")

	fmt.Println("start")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("error ocurred with %v\n", err)
		os.Exit(2)
	}

	fmt.Println("running")

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("error ocurred with %v\n", err)
	}
	fmt.Println("end")
}
