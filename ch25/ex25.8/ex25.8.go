package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())	// 컨텍스트 생성
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()	// 취소

	wg.Wait()
}