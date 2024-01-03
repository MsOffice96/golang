package main

import "log"

// 메서드 오버라이딩 지원
// 메서드 오버로딩 미지원, 이름은 같지만 매개변수가 다른 메서드 지원 x

type Item struct {
	Price    int
	Quantity int
	Name     string
}

func (t Item) Cost() float64 {
	return float64(t.Price) * float64(t.Quantity)
}

type DisCountItem struct {
	DiscountRate float64
	Item         Item
}

func (t DisCountItem) Cost() float64 {
	return t.Item.Cost() * ((100 - t.DiscountRate) * 0.01)
}

// 오버로딩과 비슷한 효과를 주기 위한 방법 , 가변 인자 표기
func display(values ...interface{}) {
	for i := 0; i < len(values); i++ {
		log.Println("value ", values[i])
	}
}
