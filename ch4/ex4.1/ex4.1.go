package main

import "fmt"

func main() {
	var a int = 10										// a 변수 선언
	var msg string = "Hello Variable"	// msg 변수 선언

	a = 20								// a값 변경
	msg = "Good Morning"	// ms값 변경
	fmt.Println(msg, a)		// msg와 a값 출력
}