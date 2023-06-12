package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n * n)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}