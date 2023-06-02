package main

import "fmt"

func main() {
	str := "Hello 월드"		 // 한글과 영문자가 섞인 문자열
	runes := []rune(str)	// []rune 타입으로 타입 변환

	fmt.Printf("len(str) = %d\n", len(str))			// string 타입 길이
	fmt.Printf("len(runes) = %d\n", len(runes))	// []rune 타입 길이
}