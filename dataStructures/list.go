package dataStructures

type List[T any] struct {
	head *elList[T]
	tail *elList[T]
}

type elList[T any] struct {
	next *elList[T]
	val  T
}
