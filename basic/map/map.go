package main

import "log"

// 슬라이스, 맵, 채널, 함수, 메서드는 참조 타입
// 참조 타입은 포인터와 마찬가지로 데이터를 가지지 않고 메모리 공간을 가리킨다
// 주소값 연산은 허용 x

func main() {
	// Map test
	test_map := map[string]int{}
	test_map["test"] = 1
	for i, v := range test_map {
		log.Println("index:", i, "val:", v)
	}

	if v, ok := test_map["test1"]; ok {
		log.Println("v:", v, "ok:", ok)
	} else {
		log.Println("v:", v, "ok:", ok)
	}

	if _, ok := test_map["test"]; !ok {
		log.Println("ok", ok)
	} else {
		log.Println("ok", ok)
	}

}
