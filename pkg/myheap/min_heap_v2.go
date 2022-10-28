package myheap

type MinHeapV2 []int

func (h MinHeapV2) Len() int {
	return len(h)
}

func (h MinHeapV2) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *MinHeapV2) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeapV2) Push(value interface{}) {
	(*h) = append((*h), value.(int))

}

func (h *MinHeapV2) Pop() interface{} {
	res := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return res
}

func (h *MinHeapV2) Peek() interface{} {
	return (*h)[0]
}
