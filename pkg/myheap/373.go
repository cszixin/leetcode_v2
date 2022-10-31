package myheap

import "container/heap"

type pair struct {
	u   int
	v   int
	sum int
}
type myheap []pair

func (h myheap) Len() int {
	return len(h)

}

func (h myheap) Less(i, j int) bool {
	return h[i].sum > h[j].sum
}

func (h *myheap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myheap) Push(value interface{}) {
	(*h) = append((*h), value.(pair))
}

func (h *myheap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &myheap{}
	heap.Init(h)
	res := make([][]int, 0)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			heap.Push(h, pair{nums1[i], nums2[j], nums1[i] + nums2[j]})
		}
	}

	i := 0
	for h.Len() > 0 {
		if h.Len() > k {
			heap.Pop(h)
		} else {
			tmp, _ := heap.Pop(h).(pair)
			res = append(res, []int{tmp.u, tmp.v})
		}
		i++
	}

	return res

}
