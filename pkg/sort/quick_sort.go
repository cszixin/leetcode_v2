package sort

func QuickSort(array []int, low, high int) {
	if low >= high {
		return
	}
	provit := partition(array, low, high)
	QuickSort(array, 0, provit-1)
	QuickSort(array, provit+1, high)
}

func swap(array []int, a, b int) {
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}

func partition(array []int, low, high int) int {
	// 基准
	start := low
	p := array[low]
	for low < high {
		for high > low && array[high] > p {
			high--
		}
		for high > low && array[low] <= p {
			low++
		}
		swap(array, low, high)
	}
	// low与high重合后，基准的位置确定
	swap(array, start, low)
	return high
}
