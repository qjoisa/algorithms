package sort

func QuickSort(array []int) []int {
	return sort(array, 0, len(array)-1)
}

func sort(arr []int, low, high int) []int {
	if low < high {
		var pivot int
		arr, pivot = partition(arr, low, high)
		arr = sort(arr, low, pivot-1)
		arr = sort(arr, pivot+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
