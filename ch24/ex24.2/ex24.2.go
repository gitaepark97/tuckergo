package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAuto(a, b int) {
	sum := 0
	for i := 0; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAuto(1, 1000000000)
	}

	wg.Wait()
}