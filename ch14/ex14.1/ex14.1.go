package main

import "fmt"

func main() {
	var a int = 500
	var p *int = &a	// a의 메모리 주소를 변수 p값으로 대입(복사)

	fmt.Printf("p의 값: %p\n", p)									// 메모리 주소값 출력
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)	// p가 가리키는 메모리의 값 출력

	*p = 100											 // p가 가리키는 메모리 공간의 값을 변경합니다.
	fmt.Printf("a의 값: %d\n", a)		// a 값의 변화 확인
}