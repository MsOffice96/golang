package main

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// i : 새로 들어온 값, j : 기존 값
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Less(i, j int) bool {
	// h[i] > h[j] : 최대 힙 -> 새로운 값(i)가 기존 값(j) 보다 큰 경우 return true & swap
	// h[i] < h[j] : 최소 힙 -> 새로운 값(i)가 기존 값(j) 보다 작은 경우 return true & swap
	return h[i] > h[j]
}

func (h *IntHeap) Push(elem interface{}) {
	*h = append(*h, elem.(int))
}

func (h *IntHeap) Pop() interface{} {
	// Interface 내부 함수에서 먼저 마지막 노드와 첫번째 노드를 바꾸고 나서
	// 내가 정의한 Pop()함수 실행 및 마지막 노드 return
	old := *h
	n := len(old)
	elem := old[n-1]
	*h = old[0 : n-1]

	return elem
}

// func main() {
// 	h := &IntHeap{9, 6, 41, 1}
// 	heap.Init(h)
// 	fmt.Println(h)

// 	heap.Push(h, 7)
// 	fmt.Println(h)

// 	for h.Len() > 0 {
// 		m := heap.Pop(h)
// 		fmt.Println("pop:", m)
// 	}
// }
