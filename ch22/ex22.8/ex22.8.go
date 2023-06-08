package main

import "fmt"

const M = 10	// 나머지 연산의 분모

func hash(d int) int {
	return d % M	// 나머지 연산
}

func main() {
	m := [M]int{}

	m[hash(23)] = 10
	m[hash(259)] = 50

	fmt.Printf("%d = %d\n", 23, m[hash(23)])
	fmt.Printf("%d = %d\n", 259, m[hash(259)])
}