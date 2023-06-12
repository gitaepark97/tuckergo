package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{ 0 }	// 0원 잔고 통장
	wg.Add(10)								// WaitGroup 객체 생성
	for i := 0; i < 10; i++ {
		go func () {
			for {
				DepositAndWithdraw(account)
			}											// 고루팅 10개 생성
		}()
	}
	wg.Wait()
}