package main

import "fmt"

type Data struct {	// Data 타입 구조체
	value int
	data 	[200]int
}

func ChangeData(arg Data) {	// 매개변수로 Data를 받습니다.
	// arg 데이터를 변경합니다.
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(data)	// 인수로 data를 넣습니다
	// data의 두 필드 출력
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}