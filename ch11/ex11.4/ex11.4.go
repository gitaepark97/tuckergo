package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {		// 3번 반복합니다.
		for j := 0; j < 5; j++ {	// 5번 반복합니다.
			fmt.Print("*")					// *을 출력합니다.
		}
		fmt.Println()							// 줄바꿈합니다.
	}
}