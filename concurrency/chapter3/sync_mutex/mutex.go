package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	mutex()
	RWmutex()
}

// Mutex (mutual exclusion) 상호배제 : 임계 영역을보호하는 방법
func mutex() {
	var lock sync.Mutex
	var cnt int

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		cnt++
		log.Printf("increment : %d\n", cnt)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		cnt--
		log.Printf("decrement: %d\n", cnt)
	}

	var aritmetic sync.WaitGroup
	for i := 0; i < 10; i++ {
		aritmetic.Add(1)
		go func() {
			defer aritmetic.Done()
			increment()
		}()
	}

	for i := 0; i < 10; i++ {
		aritmetic.Add(1)
		go func() {
			defer aritmetic.Done()
			decrement()
		}()
	}

	aritmetic.Wait()
}

// RWMutex
func RWmutex() {
	// Locker는 interface type으로 sync.Mutex, sync,RWMutex를 만족함
	// type Locker interface{
	// Lock()
	// Unlock()
	// }

	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		l.Unlock()
		time.Sleep(1)
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(cnt int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(cnt + 1) // ?
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := cnt; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")

	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()),
			test(count, &m, &m),
		)
	}
}
