package main

import "fmt"

// 상수값에 코드를 부여합니다.
const Pig int = 0
const Cow int = 1
const Chicken int = 2

// 코드값에 따라서 다른 텍스트를 출력합니다.
func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
    fmt.Println("꼬끼오")
  } else {
		fmt.Println("...")
	}
}

func main() {
	// PrintAnimal() 함수를 호출하여 동물 소리를 출력합니다.
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)
}