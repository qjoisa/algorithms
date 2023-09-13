package sort

func SelectedSort(array []int) []int {
	newArray := make([]int, len(array))
	for i := 0; i < len(newArray); i++ {
		smallest := findSmallest(array)
		newArray[i] = array[smallest]
		array = append(array[:smallest], array[smallest+1:]...)
	}
	return newArray
}

func findSmallest(array []int) int {
	min := 0
	for i := 0; i < len(array); i++ {
		if array[i] < array[min] {
			min = i
		}
	}
	return min
}
