package myheap

import (
	"container/heap"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type nodeHeap []*ListNode

func (h nodeHeap) Len() int {
	return len(h)
}
func (h nodeHeap) Less(i, j int) bool {
	//小顶堆
	return h[i].Val < h[j].Val
}

func (h *nodeHeap) Swap(i, j int) {

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *nodeHeap) Push(value interface{}) {
	(*h) = append((*h), value.(*ListNode))
}

func (h *nodeHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 常用技巧,使用的一个虚拟节点避免处理头节点等边界问题
	dummpy := &ListNode{}
	dummpy.Next = nil
	dummpy.Val = math.MinInt
	p := dummpy
	h := &nodeHeap{}
	heap.Init(h)
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}
	for h.Len() > 0 {
		tmp, _ := heap.Pop(h).(*ListNode)
		p.Next = tmp
		p = p.Next
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}

	}
	return dummpy.Next
}
