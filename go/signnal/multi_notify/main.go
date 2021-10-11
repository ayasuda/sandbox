package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	done := make(chan bool, 1)
	s1 := make(chan os.Signal, 1)
	s2 := make(chan os.Signal, 1)

	go func() {
		<-s1
		fmt.Println("got s1")
		done <- true
	}()
	go func() {
		<-s2
		fmt.Println("got s2")
		done <- true
	}()

	signal.Notify(s1, os.Interrupt)

	<-done

	signal.Notify(s2, os.Interrupt)

	<-done
}
