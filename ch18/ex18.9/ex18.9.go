package main

import "fmt"

func main() {
	slice1 := []int{ 1, 2, 3, 4, 5 }
	slice2 := make([]int, 3, 10)	// len:3, cap:10 슬라이스
	slice3 := make([]int, 10)			// len:10, cap:10 슬라이스

	cnt1 := copy(slice2, slice1)	// slice1을 slice3에 복사합니다.
	cnt2 := copy(slice3, slice1)	// slice1을 slice3에 복사합니다.

	fmt.Println(cnt1, slice2)
	fmt.Println(cnt2, slice3)
}