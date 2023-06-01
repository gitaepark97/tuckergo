package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {				// 5번 박복합니다.
		for j := 0; j < i + 1; j++ {	// 현재 i값+1만큼 반복합니다.
			fmt.Print("*")							// *을 출력합니다.
		}
		fmt.Println()									// 줄바꿈합니다.
	}
}