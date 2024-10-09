package utils

import "fmt"

type Stack struct {
	maxSize int
	stack   []int
	top     int
}

func (s *Stack) Initialize(maxSize int) {
	s.maxSize = maxSize
	s.stack = make([]int, 0, maxSize)
	s.top = -1
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == s.maxSize-1
}

func (s *Stack) Push(value int) {
	if s.IsFull() {
		fmt.Println("Stack Overflow")
		return
	}
	s.top += 1
	s.stack = append(s.stack, value)
	fmt.Printf("Succesfully Pushed")
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Printf("Stack Underflow")
		return -1
	}
	popped := s.stack[s.top]
	s.stack = s.stack[:s.top]
	s.top--
	return popped
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Printf("Stack is empty")
		return -1
	}
	return s.stack[s.top]
}


