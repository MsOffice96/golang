package main

import (
	"log"
	"sync"
	"time"
)

// Starvation Condition
// 동시 프로세스가 작업을 수행하는데 필요한 모든 리소스를 얻을수 없는 모든 상황 의미

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1 * time.Second

	greedyWorker := func() {
		defer wg.Done()

		var cnt int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			cnt++
		}
		log.Printf("Greedy Worker was able to execute %d working\n", cnt)
	}

	politeWorke := func() {
		defer wg.Done()

		var cnt int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()

			cnt++
		}

		log.Printf("Polite Worker was able to execute %d working\n", cnt)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorke()

	wg.Wait()
}
