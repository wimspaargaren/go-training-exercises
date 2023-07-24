package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx1 := context.Background()
	ctx2, c2 := context.WithCancel(ctx1)
	defer c2()

	ctx3, c3 := context.WithCancel(ctx2)
	defer c3()
	go doSth1(ctx2)

	go doSth2(ctx3)
	c3()
	time.Sleep(time.Second)
}

func doSth1(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("first goroutine return")
		return
	}
}

func doSth2(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("second goroutine return")
		return
	}
}
