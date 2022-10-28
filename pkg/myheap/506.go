package myheap

import (
	"container/heap"
	"fmt"
	"sort"
)

type Node struct {
	index int
	value int
}
type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}
func (h NodeHeap) Less(i, j int) bool {
	return h[i].value > h[j].value
}

func (h *NodeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *NodeHeap) Push(value interface{}) {
	(*h) = append((*h), value.(Node))
}
func (h *NodeHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return res
}

func findRelativeRanks(score []int) []string {
	res := make([]string, len(score))
	st := map[int]string{
		0: "Gold Medal",
		1: "Silver Medal",
		2: "Bronze Medal",
	}
	h := &NodeHeap{}
	heap.Init(h)
	for k, v := range score {
		heap.Push(h, Node{k, v})
	}
	i := 0
	for h.Len() > 0 {
		node, _ := heap.Pop(h).(Node)
		index := node.index
		txt := ""
		if i <= 2 {
			txt = st[i]
		} else {
			txt = fmt.Sprintf("%d", i+1)
		}
		res[index] = txt
		i++
	}
	return res
}

func findRelativeRanksv2(score []int) []string {
	res := make([]string, len(score))
	st := map[int]string{
		0: "Gold Medal",
		1: "Silver Medal",
		2: "Bronze Medal",
	}

	type pair struct {
		ids   int
		score int
	}
	tmp := make([]pair, 0)
	for k, v := range score {
		tmp = append(tmp, pair{k, v})
	}
	sort.Slice(tmp, func(i, j int) bool { return tmp[i].score > tmp[j].score })
	for i := 0; i < len(tmp); i++ {
		txt := ""
		if i <= 2 {
			txt = st[i]
		} else {
			txt = fmt.Sprintf("%d", i+1)
		}
		res[tmp[i].ids] = txt

	}
	return res
}
