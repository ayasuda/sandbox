package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func infLoop(ctx context.Context) {
	fmt.Println("start infLoop")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit infLoop")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("start")

	go infLoop(ctx)

	time.Sleep(1000 * time.Millisecond)
	fmt.Println("do cancel")
	cancel()

	time.Sleep(2000 * time.Millisecond)
}
