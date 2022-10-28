package myheap

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	// 定义大顶堆，用大于号
	return h[i] > h[j]
}
func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {

	//记住，最后元素是最值
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func SortByHeap(nums []int) []int {

	h := &MaxHeap{}
	heap.Init(h)
	for _, value := range nums {
		heap.Push(h, value)
	}
	size := len(nums)
	for i := 0; i < size; i++ {
		nums[i], _ = heap.Pop(h).(int)
	}
	return nums
}

// 第k小的数
//找到最小的k个数,返回这k个数中最大的
func FindKth(nums []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		//保持堆的大小的k
		heap.Push(h, nums[i])
		if i >= k {
			heap.Pop(h)
		}
	}
	//此时堆顶就是第k小的
	tmp, _ := heap.Pop(h).(int)
	return tmp
}
