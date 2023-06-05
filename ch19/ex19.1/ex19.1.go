package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {	// 일반 항수 표현
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {	// 메서드 표현
	a.balance -= amount
}

func main() {
	a := &account{ 100 }	// balance가 100인 acount 포인터 변수 생성

	withdrawFunc(a, 30)

	a.withdrawMethod(30)

	fmt.Printf("%d \n", a.balance)
}