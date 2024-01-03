package main

import (
	"log"
)

func main() {
	// Pointer
	val := 1
	var Pex *int = &val
	log.Println("val:", val)   // 1
	log.Println("&val:", &val) // 0xc000016120
	log.Println("Pex:", Pex)   // 0xc000016120
	log.Println("&Pex:", &Pex) // 0xc00000e028
	log.Println("*Pex:", *Pex) // 1

	// Struct
	rectangle := rect{
		w: 1,
		h: 2,
	}
	log.Println(rectangle)  // {1 2}
	log.Println(&rectangle) // &{1 2}

	var Prect *rect = &rectangle
	log.Println(Prect) // &{1 2} // 구조체 log 출력시 &는 주소값을 의미
	Prect.w = 2
	log.Println(*Prect) // {2 2}

	// Pointer Struct
	// new() 타입에 맞는 스페이스를 찿아서 초기화 및 주소 반환
	var PResizeRect = new(rect)
	PResizeRect.h = 1
	PResizeRect.w = 1

	Presize(PResizeRect, 1, 1)
	log.Println("PresizeReact", PResizeRect)

	// Struct
	ResizeRect := rect{
		w: 1,
		h: 1,
	}
	resize(ResizeRect, 1, 1)
	log.Println("ResizeRect", ResizeRect)

}

// Test struct
type rect struct{ w, h float32 }

func Presize(rectangle *rect, width float32, height float32) {
	rectangle.w += width
	rectangle.h += height
}

func resize(rectangle rect, width float32, height float32) {
	rectangle.w += width
	rectangle.h += height
	log.Println("func resize", rectangle)
}
