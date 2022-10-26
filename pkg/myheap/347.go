package myheap

import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	m1 := make(map[int]int, 0)
	m2 := make(map[int][]int, 0)
	H := NewMinHeap()
	for _, k := range nums {
		m1[k]++
	}

	for k, v := range m1 {
		m2[v] = append(m2[v], k)
		H.Insert(v)
	}

	s := H.GetSize() - k
	for i := 0; i < s; i++ {
		H.DelMin()
	}
	topK := make([]int, 0)
	s = H.GetSize()
	for i := 0; i < s; i++ {
		tmp := H.DelMin()
		elems := m2[tmp]
		elem := elems[0]
		elems = elems[1:]
		m2[tmp] = elems

		topK = append(topK, elem)
	}
	return topK
}

func topKFrequentV2(nums []int, k int) []int {
	m1 := make(map[int]int, 0)
	for _, v := range nums {
		m1[v]++
	}
	h := &MinHeapV2{}
	heap.Init(h)
	for key, value := range m1 {
		heap.Push(h, Node{key, value})
		if h.Len() > k {
			print(h.Len())
			heap.Pop(h)
		}
	}
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		tmp := heap.Pop(h)
		node, _ := tmp.(Node)
		res = append(res, node.key)
	}
	return res

}
