package main

import "fmt"

func main() {
	str1 := "가나다라마"	// 한글 문자열
	str2 := "abcde"			// 영문 문자열

	fmt.Printf("len(str1) = %d\n", len(str1))	// 한글 문자열 크기
	fmt.Printf("len(str2) = %d\n", len(str2))	// 영문 문자열 크기
}