package main

import "fmt"

func main() {
	var slice = []int{ 1, 2, 3 }	// 요소가 3개인 슬라이스

	slice2 := append(slice, 4)	// 요소 추가

	fmt.Println(slice)
	fmt.Println(slice2)
}