package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	done := make(chan bool, 1)
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt)

	cnt := 0

	go func() {
		for {
			<-ticker.C
			cnt += 1
			fmt.Println(cnt)
		}
	}()

	signal.Stop(sigs)

	go func() {
		// Naver reached! because Reset called
		<-sigs
		fmt.Println("bye")
		ticker.Stop()
		done <- true
	}()

	<-done
}
