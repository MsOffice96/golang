package main

import "fmt"

type Heap struct {
	list []int
}

// 최대 힙
func (h *Heap) push(v int) {
	h.list = append(h.list, v)
	idx := len(h.list) - 1

	for idx >= 0 {
		parent_idx := (idx - 1) / 2
		if parent_idx < 0 {
			break
		}
		if h.list[idx] > h.list[parent_idx] {
			h.list[idx], h.list[parent_idx] = h.list[parent_idx], h.list[idx]
			idx = parent_idx
		} else {
			break
		}
	}
}

func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}
	if len(h.list) == 1 {
		r := h.list[0]
		h.list = h.list[:len(h.list)-1]
		fmt.Println("extract root node : ", h.list)
		return r
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	h.list[0] = last

	idx := 0

	for idx < len(h.list) {
		if idx >= len(h.list) {
			break
		}
		swapIdx := -1
		leftIdx := idx*2 + 1

		if leftIdx >= len(h.list) {
			break
		}

		if h.list[leftIdx] > h.list[idx] {
			swapIdx = leftIdx
		}

		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] > h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}

		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx

	}
	return top
}

func (h *Heap) print() {
	fmt.Println(h.list)
}

func main() {
	fmt.Println("Priority Queue")
	h := &Heap{}

	for i := 0; i < 5; i++ {
		h.push(i)
		h.print()
	}

	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	// for j := 0; j < 2000; j++ {
	// 	fmt.Println(h.Pop())
	// }

}
