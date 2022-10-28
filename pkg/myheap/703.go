package myheap

import (
	"container/heap"
)

type KthLargest struct {
	h *MinHeapV2
	k int
}

func Constructor(k int, nums []int) KthLargest {
	h := MinHeapV2(nums)
	heap.Init(&h)
	for h.Len() > k {
		heap.Pop(&h)
	}
	st := KthLargest{&h, k}
	return st
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)
	for this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	res, _ := this.h.Peek().(int)
	return res
}
