package main

import "fmt"

func main() {
	str := "Hello 월드!"			 // 한영 문자가 섞인 문자여
	for _, v := range str {		// range를 이용한 순회
		fmt.Printf("타입:%T 값:%d 문자:%c\n", v, v, v)	// 출력
	}
}