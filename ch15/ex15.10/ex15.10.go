package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2	// str1, " ", str2를 잇습니다.
	fmt.Println(str3)

	str1 += " " + str2				// str1에 " " + str2 문자열을 붙입니다.
	fmt.Println(str1)
}