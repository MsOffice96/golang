package main

import (
	"fmt"
	"sync"
	"time"
)

// Cond
// 고루틴들이 대기하거나, 어떤 이벤트의 발생을 알리는 집결지점
// 위에서의 이벤트는 두개이상의 고루틴 사이에서 어떠한 사실이 발생했다는 사실외에는 아무런 정보를 전달하지 않는 임의의 신호를 말함

func main() {
	Cond()
}

func Cond() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromeQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Printf("Removed from queue / Queue Length: %d\n", len(queue))
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Printf("Adding to queue / Queue Length: %d\n", len(queue))
		queue = append(queue, struct{}{})
		go removeFromeQueue(1 * time.Second)
		c.L.Unlock()
	}
}
