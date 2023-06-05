package main

import "fmt"

func changeArray(array2 [5]int) {	// 배열을 받아서 세 번째 값 변경
	array2[2] = 200
}

func changeSlice(slice2 []int) {	// 슬라이스를 받아서 세 번째 값 변경
	slice2[2] = 200
}

func main() {
	array := [5]int{ 1, 2, 3, 4, 5 }
	slice := []int{ 1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
}