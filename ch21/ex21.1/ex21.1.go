package main

import "fmt"

func sum(nums ...int) int {	// 가변 인수를 받는 함수
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)	// nums 타입 출력
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}