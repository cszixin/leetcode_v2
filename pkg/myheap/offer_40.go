package myheap

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
