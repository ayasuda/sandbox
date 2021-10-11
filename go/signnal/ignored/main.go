package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	tm := time.NewTimer(10 * time.Second)
	ti := time.NewTicker(1 * time.Second)
	sigs := make(chan os.Signal, 1)

	fmt.Println("ignore Ctrl + C")
	signal.Ignore(os.Interrupt)

	go func() {
		for {
			<-ti.C
			if signal.Ignored(os.Interrupt) {
				fmt.Println("ignoring Ctrl + C")
			} else {
				fmt.Println("wait Ctrl + C")
			}
		}
	}()

	go func() {
		<-tm.C
		tm.Stop()

		fmt.Println("start to accept Ctrl + c")
		signal.Notify(sigs, os.Interrupt)
	}()

	<-sigs
	ti.Stop()
	fmt.Println("bye")
}
