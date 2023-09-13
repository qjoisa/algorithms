package search

func BinarySearch(array []int, value int) int {
	max := len(array) - 1
	min := 0
	for min <= max {
		mid := (max + min) / 2
		if array[mid] == value {
			return mid
		} else if array[mid] < value {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return 0
}
