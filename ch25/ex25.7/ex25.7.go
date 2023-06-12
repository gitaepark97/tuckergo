package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body 	string
	Tire 	string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func MakeBody(tireCh chan *Car) {	// 차체 생상
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <- tick:
			// Make body
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {	// 바퀴 설치
	for car := range tireCh {
		// install a tire
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) {	// 도색
	for car := range paintCh {
		// paint
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Since(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("Start Factory")

	wg.Add(3)
	go MakeBody(tireCh)	// 고루틴 생성
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}