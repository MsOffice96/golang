package main

import (
	"log"
	"sync"
	"time"
)

// WaitGroup
// 동시에 수행된 연산의 결과를 신경 쓰지 않거나
// 결과를 수집할 다른 방법이 있는 경우 동시에 수행될 연산 집합을 기다릴 때 유용함.
// func main() {

// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		log.Println("1")
// 		time.Sleep(time.Second * 1)
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		log.Println("2")
// 		time.Sleep(time.Second * 1)
// 	}()

// 	wg.Wait()
// }

// sync.WaitGroup의 Add는 가급적 고루틴의 외부에서 수행하도록한다.
// 언제 고루틴이 스케줄링이 될지 모르는 상황에서
// sync.WaitGroup의 Wait()에 고루틴의 실행 전 즉 Add()가 시작되기 전에 도달 가능성 때문이다.
// 밑의 예시는 WaitGroup을 Parameter로 받아 고루틴에서 Add를 하며
// main 함수의 wait에 먼저 도달하여 고루틴이 실행되지 않는 예시이다.
// func WaitGroup_Error(wg *sync.WaitGroup, i int) {
// 	// if wg != nil {
// 	wg.Add(1)
// 	defer wg.Done()
// 	// }
// 	log.Printf("WaitGroup Error, goroutine index %d", i)
// 	time.Sleep(time.Second * 1)
// }

func main() {
	var wg sync.WaitGroup

	WaitGroup_Error := func(wg *sync.WaitGroup, i int) {
		wg.Add(1)
		defer wg.Done()
		log.Printf("WaitGroup Error, goroutine index %d", i)
		time.Sleep(time.Second * 1)
	}

	working := 10
	// wg.Add(working)
	for i := 0; i < working; i++ {

		go WaitGroup_Error(&wg, i)
	}

	wg.Wait()
}
