package myheap

type Node struct {
	key   int
	value int
}

type MinHeapV2 []Node

func (h MinHeapV2) Len() int {
	return len(h)
}
func (h MinHeapV2) Less(i, j int) bool {
	return h[i].value < h[j].value
}
func (h *MinHeapV2) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinHeapV2) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *MinHeapV2) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
