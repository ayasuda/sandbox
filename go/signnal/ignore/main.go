package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	t := time.NewTimer(10 * time.Second)
	sigs := make(chan os.Signal, 1)

	fmt.Println("ignore Ctrl + C")
	signal.Ignore(os.Interrupt)

	go func() {
		<-t.C
		t.Stop()

		fmt.Println("start to accept Ctrl + c")
		signal.Notify(sigs, os.Interrupt)
	}()

	<-sigs
	fmt.Println("bye")
}
