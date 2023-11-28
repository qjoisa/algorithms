package dataStructures

import "errors"

type Queue[T any] struct {
	current, last *elQueue[T]
}

type elQueue[T any] struct {
	next *elQueue[T]
	val  T
}

func (q *Queue[T]) Push(v T) {
	if q.current == nil {
		q.last = &elQueue[T]{val: v}
		q.current = q.last
	} else {
		q.last.next = &elQueue[T]{val: v}
		q.last = q.last.next
	}
}

func (q *Queue[T]) Pop() (v T, err error) {
	if q.current == nil {
		return v, errors.New("Queue is empty")
	}
	v = q.current.val
	q.current = q.current.next
	return v, nil
}

func (q *Queue[T]) IsEmpty() bool {
	if q.current == nil {
		return true
	} else {
		return false
	}
}
