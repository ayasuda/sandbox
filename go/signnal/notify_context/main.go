package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	parent, cancel := context.WithCancel(context.Background())

	ctx, stop := signal.NotifyContext(parent, os.Interrupt)
	defer stop()

	go func(ctx context.Context) {
		// 親から context 受け取ってるけど特にすることないので無視
		for {
			var in string
			fmt.Scanln(&in)
			fmt.Println(in)
			switch in {
			case "cancel":
				fmt.Println("canceled")
				cancel()
			case "exit":
				fmt.Println("exited")
				stop()
			default:
				fmt.Println(in)
			}
		}
	}(ctx)

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("time out")
	case <-ctx.Done():
		stop()
		fmt.Println("interuppted")
	}
}
