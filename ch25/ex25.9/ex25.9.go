package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Key string

func square(ctx context.Context) {
	if v := ctx.Value(Key("number")); v != nil {
		n := v.(int)
		fmt.Printf("Sqaure:%d", n * n)
	}
	wg.Done()
}

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), Key("number"), 9)
	// 컨텍스트에 값을 추가한다.
	go square(ctx)

	wg.Wait()
}