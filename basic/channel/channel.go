package main

import (
	"log"
	"strconv"
	"time"
)

// 타임아웃 : 시간이 오래 걸리는 작업에 타임아웃 처리하기
func process(quit <-chan struct{}) chan string {
	done := make(chan string) // Process 완료 신호를 위한 channel

	go func() {
		go func() {
			time.Sleep(2 * time.Second) // heavy job

			done <- "Complete"
		}()
		<-quit // parameter로 받아온 channel(quit) 의 main 종료 신호시에 종료
		return
	}()
	return done
}

// 공유 메모리 : 채널을 사용하여 여러 고루틴의 공유 메모리 접근 제어 ()
const (
	set = iota
	get
	remove
	count
)

type ShareMap struct {
	memory map[string]interface{}
	cmd    chan Command
}

type Command struct {
	action int
	key    string
	value  interface{}
	result chan interface{}
}

func NewSharMap() ShareMap {
	sm := ShareMap{
		memory: make(map[string]interface{}),
		cmd:    make(chan Command),
	}
	go sm.run()
	return sm
}

func (sm *ShareMap) run() { // Create Struct & start run() goroutine
	for cmd := range sm.cmd {
		switch cmd.action {
		case set:
			sm.memory[cmd.key] = cmd.value
			val, ok := sm.memory[cmd.key]
			log.Println("case", val, ok)
			// cmd의 result의 chan은 method에서 생성하여 전달 받음
			cmd.result <- []interface{}{val, ok}
		case get:
			val, ok := sm.memory[cmd.key]
			cmd.result <- []interface{}{val, ok}
		case remove:
			_, ok := sm.memory[cmd.key]
			if ok {
				delete(sm.memory, cmd.key)
				cmd.result <- []interface{}{cmd.key, true}
			} else {
				cmd.result <- []interface{}{cmd.key, false}
			}
		case count:
			cnt := len(sm.memory)
			cmd.result <- []interface{}{cnt}
		default:
			log.Println("Working")
		}
	}
}

func (sm *ShareMap) set(k string, val interface{}) (interface{}, bool) {
	callback := make(chan interface{})
	sm.cmd <- Command{action: set, key: k, value: val, result: callback}

	cmd_result := (<-callback).([]interface{})

	// log.Println("set result:", cmd_result[0].(bool), cmd_result[1])
	return cmd_result[0], cmd_result[1].(bool)
	// return cmd_result[0].(bool), cmd_result[1]
}

func (sm *ShareMap) get(k string) (interface{}, bool) {
	callback := make(chan interface{})
	sm.cmd <- Command{action: get, key: k, result: callback}
	cmd_result := (<-callback).([]interface{})
	return cmd_result[0], cmd_result[1].(bool)
	// callback := make(chan interface{})
	// sm.cmd <- Command{action: get, key: k}
	// cmd_result := (<-callback).([2]interface{})
	// return cmd_result[0].(bool), cmd_result[1]
}

func (sm *ShareMap) remove(key string) (interface{}, bool) {
	callback := make(chan interface{})
	sm.cmd <- Command{action: remove, key: key, result: callback}
	cmd_result := (<-callback).([]interface{})
	return cmd_result[0], cmd_result[1].(bool)
}

func (sm *ShareMap) count() (cnt int) {
	callback := make(chan interface{})
	sm.cmd <- Command{action: count, result: callback}
	cmd_result := (<-callback).([]interface{})
	return cmd_result[0].(int)
}

// 파이프라인 : 여러 고루틴을 파이프라인 형태로 연결

// 파이프라인 - Handler를 사용하지 않을 경우
const (
	JOB_COUNT = 500
	BUF_SIZE  = 500
)

type Job struct{ name, log string }

func (j Job) String() string {
	return "job name: " + j.name + "\n[log]\n" + j.log
}

func prepare() <-chan Job {
	out := make(chan Job, BUF_SIZE)
	go func() {
		//prepare job
		for i := 0; i < JOB_COUNT; i++ {
			// time.Sleep(time.Millisecond * 1000)
			out <- Job{name: strconv.Itoa(i)}
		}
		close(out)
	}()
	return out
}

func doFirst(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			job.log += "First stage\n"
			out <- job
		}
		close(out)
	}()
	return out
}

func doSecond(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			job.log += "second stage \n"
			out <- job
		}
		close(out)
	}()
	return out
}

func doThird(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			job.log += "third stage \n"
			out <- job
		}
		close(out)
	}()
	return out
}

func doLast(in <-chan Job) <-chan Job {
	out := make(chan Job, cap(in))
	go func() {
		for job := range in {
			job.log += "last stage\n"
			out <- job
		}
		close(out)
	}()
	return out
}

/*
	맵리듀스 : 고루틴을 사용하여 맵리듀스 패턴 구현
	맵(map): 흩어져 있는 데이터 조각을 종류별로 모으는 단계
	리듀스(reduce): 맵 작업으로 생성된 데이터를 모두 취합하여 원하는 형태의 최종정보를 추출하는 단계
*/
