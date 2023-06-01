package main

import "fmt"

const Y int = 3	// 상수

func main() {
x := 5				// 변수
	// a := [x]int{1, 2, 3, 4, 5}

	b := [Y]int{1, 2, 3}

	var c [10]int

	fmt.Println(x)
	fmt.Println(b)
	fmt.Println(c)
}