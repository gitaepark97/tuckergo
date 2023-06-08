package main

import "fmt"

type Product struct {
	Name 	string
	Price int
}

func main() {
	m := make(map[int]Product)	// 맵 생성
	// 요소 추가
	m[16] = Product{ "볼펜", 500 }
	m[46] = Product{ "볼펜", 200 }
	m[78] = Product{ "볼펜", 1000 }
	m[345] = Product{ "볼펜", 3000 }
	m[897] = Product{ "볼펜", 500 }

	for k, v := range m {	// 맵 순회
		fmt.Println(k, v)
	}
}