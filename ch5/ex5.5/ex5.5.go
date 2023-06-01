package main

import "fmt"

func main() {
	var a int										// 값을 저장할 변수
	var b int

	n, err := fmt.Scan(&a, &b)	// 입력 두개 받기
	if err != nil {							// 에러가 발생하면 에러 코드 출력
		fmt.Println(n, err)
	} else {										// 정상 입력되면 입력값 출력
		fmt.Println(n, a, b)
	}
}