package dataStructures

import "errors"

type Stack[T any] struct {
	node *elStack[T]
}

type elStack[T any] struct {
	prev *elStack[T]
	val  T
}

func (s *Stack[T]) Push(v T) {
	if s.node == nil {
		s.node = &elStack[T]{val: v}
	} else {
		s.node = &elStack[T]{prev: s.node, val: v}
	}
}

func (s *Stack[T]) Pop() (v T, err error) {
	if s.node == nil {
		return v, errors.New("Stack is empty")
	} else {
		v = s.node.val
		s.node = s.node.prev
		return v, nil
	}
}

func (s *Stack[T]) IsEmpty() bool {
	if s.node == nil {
		return true
	} else {
		return false
	}
}
