package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		// break	// break 사용
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
	case 4:	// break 사용 안함
		fmt.Println("a == 4")
	default:
		fmt.Println("a > 4")
	}
}