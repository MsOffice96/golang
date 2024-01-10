package main

import (
	"bytes"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// Livelock Condition
// 프로그램들이 활동적으로 동시에 연산을 수행하고는 있지만, 연산들이 실제로 프로그램의 상태를 진행시키는 데 아무런 영향을 주지 못하는 의미없는 연산상태를 의미

func main() {
	cadence := sync.NewCond(&sync.Mutex{})

	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
			// log.Println("tick")
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(tryName string, dirName string, dir *int32, out *bytes.Buffer) bool {
		log.Printf("%s : out %v", tryName, dirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			log.Printf("out . Success!")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32

	tryLeft := func(tryName string, out *bytes.Buffer) bool {
		log.Printf("%s try left\n", tryName)
		return tryDir(tryName, "left", &left, out)
	}

	tryRight := func(tryName string, out *bytes.Buffer) bool {
		log.Printf("%s try Right\n", tryName)
		return tryDir(tryName, "right", &right, out)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() {
			log.Println(name + out.String())
		}()

		defer walking.Done()

		log.Printf("%s : %s is trying to scoot: ", out.String(), name)
		for i := 0; i < 5; i++ {
			if tryLeft(name, &out) || tryRight(name, &out) {
				return
			}
		}
		log.Printf("%s : %v tosses her hands up in exasperation! ", &out, name)
	}

	var peopleInHallWay sync.WaitGroup
	peopleInHallWay.Add(2)

	go walk(&peopleInHallWay, "A")
	go walk(&peopleInHallWay, "B")
	peopleInHallWay.Wait()

}
