package main

import (
	"fmt"
	"log"
	"time"
)

// packaeg Errors
// type error interface {
// 	Error() string
// }

// func New(txt string) error {
// 	return &errorString{txt}
// }

// type errorString struct {
// 	err string
// }

// func (e *errorString) Error() string {
// 	return e.err
// }

// 1. Error 발생 방법
// errors.New("Not Found Error")

// 2. fmt 패키지 사용
// func Errorf(format string, a ...interface{}) error {
// 	return errors.New(fmt.Sprintf(format, a...))
// }

type sqrtError struct {
	time    time.Time
	value   float64
	message string
}

// package Errors 는 interface Error()를 포함함으로 Error()를 구현하는 경우 error type으로 반환 가능
func (sqrterror sqrtError) Error() string {
	return fmt.Sprintf("error time: %v", sqrterror.time, "value: %s", sqrterror.value, "message: %g", sqrterror.message)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{time: time.Now(), value: f, message: "minus"}
	}
	return f, nil
}

// 패니킹(panicking)
// 함수안에서 panic() 발생시 -> 현재 함수의 실행을 중지 및 종료 -> 모든 defer 구문 실행
// -> 자신을 호출한 함수로 패닉에 대한 제어권을 넘김 -> 함수 호출 스택의 상위 레벨로 올라가며 계속 이어져 main 함수에서 프로그램 종료 및 에러 상화을 출력

// recover() 함수는 패닉으로 부터 프로그램 제어권을 다시 얻어 종료 절차를 중지 후 프로그램을 계속 진행 가능
// 반드시 defer 안에서 사용해야 되며 패닉 내부의 상황을 error 값으로 복원 가능
// 패닉을 복원 후 패니킹 상황이 종료 함수 반환 타입의 제로값이 반환

func divid(a int, b int) int {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	return a / b
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
