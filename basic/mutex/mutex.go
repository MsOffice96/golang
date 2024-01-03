package main

import (
	"log"
	"sync"
	"sync/atomic"
)

type sync_dis interface {
	display()
}

func syn_display(s sync_dis) {
	s.display()
}

const initValue = -500

/* sync.Mutex (Mutex는 구조체)
Lock() : 다른 GoRoutine의 해당 메모리 접근 방지
UnLoock() : 해당 메모리의 값 변경 후 다른 GoRoutine들 접근 허용
*/

type counter struct {
	i    int
	mu   sync.Mutex
	once sync.Once
}

func (c *counter) increment() {
	c.once.Do(func() {
		log.Println("struct c init Value")
		c.i = initValue
	})
	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

func (c *counter) display() {
	log.Println("count num ", c.i)
}

func (c *counter) wait_group_increment() {
	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

/* Sync.RWMutex : 읽기 동작과 쓰기 동작을 나누어 잠금 처리
읽기 잠금 : 읽기 작업에 한해서 공유 데이터가 변하지 않음을 보장, 읽기 잠금으로 처리 경우 다른 고루틴에서 쓰기x, 읽기 o
쓰기 잠금 : 쓰기 잠금 처리 경우 , 다른 고루틴에서 쓰기x, 읽기 x

Lock() : 쓰기 잠금
Unlock() : 쓰기 잠금 해제

RLock() : 읽기 잠금
RUnlock() : 읽기 잠금 해제
*/

/* Sync.Once
func (o *Once)Do (f func())
특정 함수를 한번만 수행해야 될때 사용
한번만 수행해야 하는 함수를 Do()메서드의 매개변수로 전달하여 실행하면 여러 고루틴에서
실행하여도 해도 해당 함수는 한번만 수행됨
*/

/*
	Sync.atomic : 해당 패키지에서 제공하는 함수를 사용하는 경우 CPU에서 시분할을 하지 않고 한번에 처리 하도록 제어

Function List
Add T : 특정 포인터 변수에 값을 더함
CompareAndSwap T : 특정 포인터 변수의 값을 주어진 값과 비교하여 같으면 새로운 값으로 대처함
Load T : 특정 포인터의 변수의 값을 가져옴
Store T : 특정 포인터의 변수에 값을 저장함
Swap T : 특정 포인터 변수에 새로운 값을 저장하고 이전 값을 가져옴
*/
type atomic_s struct {
	i int64
}

func (c *atomic_s) atomic_increment() {
	atomic.AddInt64(&c.i, 1)
}

func (c *atomic_s) display() {
	log.Println("atomic_s", c.i)
}
