package main

import "fmt"

func main() {
	m := make(map[int]int)	// 맵 생성
	m[1] = 1
	m[2] = 2
	m[3] = 3

	delete(m, 3)			// 요소 삭제
	delete(m, 4)			// 없는 요소 삭제 시도
	fmt.Println(m[3])	// 삭제된 요솟값 출력
	fmt.Println(m[1])	// 존재하는 요솟값 출력
}