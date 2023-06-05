package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))	// 시간 값을 랜덤 시드로 설정

	n := rng.Intn(100)
	fmt.Println(n)
}