package main

import "fmt"

type node struct {
	val  any
	next *node
}

type Stack struct {
	top *node
	len int
}

func New() *Stack {
	return &Stack{top: nil, len: 0}
}

func (s *Stack) Push(v any) {
	n := &node{val: v, next: nil}
	if s.top != nil {
		n.next = s.top
	}
	s.top = n
	s.len++
}

func (s *Stack) Pop() any {
	if s.len == 0 {
		return nil
	}
	v := s.top.val
	s.top = s.top.next
	s.len--
	return v
}

func (s *Stack) Len() int {
	return s.len
}

func (s *Stack) Peek() any {
	if s.len == 0 {
		return nil
	}
	return s.top.val
}

func (s *Stack) Print() {
	n := s.top
	for n != nil {
		fmt.Printf("%v, ", n.val)
		n = n.next
	}
	fmt.Println()
}
