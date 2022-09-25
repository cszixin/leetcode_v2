package sort

func QuickSort(array []int, low, hight int) {
	if low >= hight {
		return
	}
	provit := partition(array, low, hight)
	QuickSort(array, 0, provit-1)
	QuickSort(array, provit+1, hight)
}

func swap(array []int, a, b int) {
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}

func partition(array []int, low, hight int) int {
	// 基准
	start := low
	p := array[low]
	for low < hight {
		for hight > low && array[hight] > p {
			hight--
		}
		for hight > low && array[low] <= p {
			low++
		}
		swap(array, low, hight)
	}
	swap(array, start, low)
	return hight
}
