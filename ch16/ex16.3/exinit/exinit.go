package exinit

import "fmt"

var (
	a = c + b	// a값은 c와 b가 초기화된 다음 초기화됩니다.
	b = f()		// b의 값은 4가 됩니다.
	c = f()		// c값은 5입니다.
	d = 3			// d값은 초기화가 끝난 뒤 6이 됩니다.
)

func init() {
	d++
	fmt.Println("init function", d)
}

func f() int {
	d++
	fmt.Println("f() d:", d)

	return d
}

func PrintD() {
	fmt.Println("d:", d)
}