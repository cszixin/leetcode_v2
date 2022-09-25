package main

import (
	"fmt"
	"leetcode/pkg/sort"
)

func main() {
	array := []int{1, 0, 0, 1}
	sort.QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
