package main

import "log"

/* interface 는 객체의 동작 표현
각 타입이 실제로 내부에 어떻게 구현되어 있는지 x, 단순한 동작 방식 표현
golang : 덕 타이핑 방식, 객체의 변수나 메서드의 집합이 객체의 타입을 결정
*/

func des(s shapper) {
	log.Println("area: ", s.area())
}

type shapper interface {
	area() float64
}

type rectangle_inter struct {
	width  float64
	height float64
}

func (rect rectangle_inter) area() float64 {
	return rect.width * rect.height * 0.5
}

type circle struct {
	rad float64
}

func (c circle) area() float64 {
	return c.rad * c.rad * 3.14
}

/* interface 이름은 메서드 이름에 "er" or "r"을 붙여서 짓는다
interface는 메서드 1~3 개 사이로 정의
*/

/* 제너릭 컬렉션
타입을 interface로 지정시 인터페이스를 충족하는 타입 값은
어떤 값이라도 배열이나 슬라이스에 담을수 있음
*/

type Figures []shapper

func (f Figures) area() (c float64) {
	for _, v := range f {
		c += v.area()
	}
	return c
}
