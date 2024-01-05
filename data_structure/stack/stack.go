package main

import (
	"errors"
	"log"
)

type Stack struct {
	s    []int
	size int
}

func NewStack() *Stack {
	stack := &Stack{}
	return stack
}

func (stack *Stack) Push(v int) {
	stack.s = append(stack.s, v)
	stack.size += 1
}

func (stack *Stack) Pop() (int, error) {
	if stack.size == 0 {
		return -1, errors.New("empty stack")
	}
	item, items := stack.s[stack.size-1], stack.s[:stack.size-1]
	stack.s = items
	stack.size -= 1
	return item, nil
}

func main() {
	log.Println("Stack")
	s := NewStack()

	if _, err := s.Pop(); err != nil {
		log.Println(err)
	}

	for i := 0; i < 10; i++ {
		s.Push(i)
		log.Println(s.s)
	}

	for {
		if v, err := s.Pop(); err != nil {
			log.Fatalln(err)
		} else {
			log.Println("Pop: ", v)
		}

	}

}
