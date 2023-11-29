package dataStructures

import "errors"

type Stack[T comparable] struct {
	node *elStack[T]
}

type elStack[T comparable] struct {
	prev *elStack[T]
	val  T
}

func (s *Stack[T]) Push(v ...T) {
	if len(v) == 0 {
		return
	}
	var i int
	if s.IsEmpty() {
		s.node = &elStack[T]{val: v[0]}
		i = 1
	}
	for ; i < len(v); i++ {
		s.node = &elStack[T]{prev: s.node, val: v[i]}
	}
}

func (s *Stack[T]) Pop() (v T, err error) {
	if s.IsEmpty() {
		return v, errors.New("Stack is empty")
	} else {
		v = s.node.val
		s.node = s.node.prev
		return v, nil
	}
}

func (s *Stack[T]) Peek() (T, bool) {
	if !s.IsEmpty() {
		return s.node.val, true
	} else {
		var v T
		return v, false
	}
}

func (s *Stack[T]) Contains(v T) bool {
	if s.IsEmpty() {
		return false
	}
	for elem := s.node; elem != nil; elem = elem.prev {
		if elem.val == v {
			return true
		}
	}
	return false
}

func (s *Stack[T]) Count() int {
	if s.IsEmpty() {
		return 0
	}
	var i int
	for elem := s.node; elem != nil; elem = elem.prev {
		i++
	}
	return i
}

func (s *Stack[T]) IsEmpty() bool {
	if s.node == nil {
		return true
	} else {
		return false
	}
}

func (s *Stack[T]) Clear() {
	s.node = nil
}
