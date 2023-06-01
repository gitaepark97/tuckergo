package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	a, b = b, a	// a와 b값을 서로 바꿉니다.
	
	fmt.Println(a, b)
}