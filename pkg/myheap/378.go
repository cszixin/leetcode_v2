package myheap

import "container/heap"

func kthSmallest(matrix [][]int, k int) int {
	h := &MinHeapV2{}
	heap.Init(h)
	k = len(matrix)*len(matrix[0]) - k + 1
	size := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			size++
			heap.Push(h, matrix[i][j])
			if size > k {
				heap.Pop(h)

			}

		}

	}
	res, _ := h.Peek().(int)
	return res
}
