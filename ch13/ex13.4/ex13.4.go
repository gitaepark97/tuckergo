package main

import "fmt"

type User struct {
	Name 	string
	ID 		string
	Age 	int
	Level int						// User의 Level 필드
}

type VIPUser struct {	// VIP 고객용 구조체
	User		
	Price int
	Level int						// VIPUser의 Level 필드
}

func main() {
	user := User{ "송하나", "hana", 23, 10 }
	vip := VIPUser{
		User{ "화랑", "hwarang", 40, 10 },
		3,
		250,	// 여러 줄로 초기화할 때는 제일 마지막 값 뒤에 꼭 쉼표를 달아주세요.
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d VIP 유저 레벨: %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,			// VIPUser의 Level
		vip.User.Level,	// 포함된 구조체명을 쓰고 접근
	)
}