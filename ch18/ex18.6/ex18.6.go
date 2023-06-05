package main

import "fmt"

func main() {
	slice1 := []int{ 1, 2, 3 }	// len:3 cap:3 슬라이스 생성

	slice2 := append(slice1, 4, 5)	// append() 함수로 요소 추가

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100	// slice2까지 바뀝니다.

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	
	slice1 = append(slice1, 500)	// slice1 요소값 변경

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}