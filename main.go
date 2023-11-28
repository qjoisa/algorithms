package main

import (
	"algorithms/dataStructures"
	"fmt"
	"log"
)

func main() {
	//unsortedArray := []int{1, 4, -2, 5, 7, -7, 9, -3, -12, 3, -9}
	//sortedArray := []int{-9, -7, -6, -3, -1, 2, 3, 5, 7, 8, 9, 11, 17, 21}
	//index := search.BinarySearch(sortedArray, 11)
	//fmt.Println(sortedArray[index])
	stack := dataStructures.Queue[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	v, err := stack.Pop()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}
