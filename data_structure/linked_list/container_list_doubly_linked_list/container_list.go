package main

import (
	"container/list"
	"fmt"
)

// container/list 를 사용한 Doubly Linked List

func main() {
	double_linked_list := list.New()

	first := double_linked_list.PushFront(1)
	double_linked_list.PushBack(2)

	final := double_linked_list.PushBack(3)

	for e := double_linked_list.Front(); e != nil; e = e.Next() {
		fmt.Printf("push : %v\n", e.Value)
	}

	double_linked_list.Remove(first)
	double_linked_list.Remove(final)

	for e := double_linked_list.Front(); e != nil; e = e.Next() {
		fmt.Printf("After remove: %v\n", e.Value)
	}

}
