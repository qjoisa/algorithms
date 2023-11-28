package dataStructures

import (
	"errors"
)

type Queue[T comparable] struct {
	current, last *elQueue[T]
}

type elQueue[T comparable] struct {
	next *elQueue[T]
	val  T
}

func (q *Queue[T]) Enqueue(v ...T) {
	if len(v) == 0 {
		return
	}
	var i int
	if q.IsEmpty() {
		i = 1
		q.last = &elQueue[T]{val: v[0]}
		q.current = q.last
	}
	for ; i < len(v); i++ {
		q.last.next = &elQueue[T]{val: v[i]}
		q.last = q.last.next
	}
}

func (q *Queue[T]) Dequeue() (v T, err error) {
	if q.IsEmpty() {
		return v, errors.New("Queue is empty")
	}
	v = q.current.val
	q.current = q.current.next
	return v, nil
}

func (q *Queue[T]) Peek() (T, bool) {
	if !q.IsEmpty() {
		return q.current.val, true
	} else {
		var v T
		return v, false
	}
}

func (q *Queue[T]) Contains(v T) bool {
	if q.IsEmpty() {
		return false
	}
	for elem := q.current; elem != nil; elem = elem.next {
		if elem.val == v {
			return true
		}
	}
	return false
}

func (q *Queue[T]) Count() int {
	var i int
	for elem := q.current; elem != nil; elem = elem.next {
		i++
	}
	return i
}

func (q *Queue[T]) IsEmpty() bool {
	if q.current == nil {
		return true
	} else {
		return false
	}
}

func (q *Queue[T]) Clear() {
	q.current = nil
	q.last = nil
}
