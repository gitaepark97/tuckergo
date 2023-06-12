package main

import "fmt"

func main() {
	ch := make(chan int)	// 크기 0인 채널 생성

	ch <- 9
	fmt.Println("Never print")
}