package main

import (
	"log"
	"time"
)

func long() {
	log.Println("long start", time.Now())
	time.Sleep(3 * time.Second)
	log.Println("long end", time.Now())
}

func short() {
	log.Println("short start", time.Now())
	time.Sleep(1 * time.Second)
	log.Println("short end", time.Now())
}

/* Goroutine
goroutine 실행 중에도 main 함수 종료시에는 프로그램 종료
Go에서 제안하는 방법: 메인 함수에서 goroutine의 종료 상황을 확인 할 수 있는
채널 생성, 만든 채널을 통해 종료 메시지를 대기
*/

func long_chan(done chan bool) {
	log.Println("long start", time.Now())
	time.Sleep(3 * time.Second)
	log.Println("long end", time.Now())
	done <- true
}

func short_chan(done chan bool) {
	log.Println("short start", time.Now())
	time.Sleep(1 * time.Second)
	log.Println("short end", time.Now())
	done <- true
}

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		log.Println("fibonacci for")
		// log.Println("fibonacci x:", x, "y:", y)
		select {
		case c <- x:
			log.Println("fibonacci x:", x, "y:", y)
			time.Sleep(time.Millisecond * 2000)
			x, y = y, x+y
		case <-quit:
			log.Println("fibonacci quit")
			return
		default:
			log.Println("fibonacci x:", x, "y:", y)
			log.Println("no signal")
		}

	}
}
