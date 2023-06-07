package main

import "fmt"

type Stringer interface {	// Stringer 인터페이스 선언
	String() string
}

type Student struct {
	Name 	string
	Age 	int
}

func (s Student) String() string {	// Student의 String() 메서드
	// 문자열 만들기
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name)
}

func main() {
	student := Student{ "철수", 12 }	 // Student 타입
	var stringer Stringer = student		// Stringer 타입

	fmt.Printf("%s\n", stringer.String())	// stringer의 String() 메서드 호출
}