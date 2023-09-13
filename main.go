package main

import (
	"algorithms/sort"
	"fmt"
)

func main() {
	unsortedArray := []int{1, 4, -2, 5, 7, -7, 9, -3, -12, 3, -9}
	//sortedArray := []int{-9, -7, -6, -3, -1, 2, 3, 5, 7, 8, 9, 11, 17, 21}
	//index := search.BinarySearch(sortedArray, 11)
	//fmt.Println(sortedArray[index])
	fmt.Println(sort.BubleSort(unsortedArray))
}
