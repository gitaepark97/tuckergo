package main

import (
	"github.com/gitaepark/tuckergo/ch20/fedex"
	koreapost "github.com/gitaepark/tuckergo/ch20/koreaPost"
)

// Sender 인터페이스를 만들었습니다.
type Sender interface {
	Send(parcel string)
}

// Sender 인터페이스를 입력으로 받습니다.

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	// 우체국 전송 객체, Fedex 전송 객체 모두 SendBook 인수로 사용할 수 있습니다.
	// 우체국 전송 객체를 만듭니다.
	koreaPostSender := &koreapost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)

	// Fedex 전송 객체를 만듭니다.
	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}