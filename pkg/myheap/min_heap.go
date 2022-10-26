package myheap

import (
	"fmt"
	"math"
)

type MinHeap struct {
	data []int
}

//构造小顶堆
func NewMinHeap() *MinHeap {
	return &MinHeap{data: []int{math.MinInt}}
}

func (H *MinHeap) swap(i, j int) {
	H.data[i], H.data[j] = H.data[j], H.data[i]
}

func (H *MinHeap) Insert(value int) {
	H.data = append(H.data, value)
	i := len(H.data) - 1
	//父比子大,上浮
	for H.data[i/2] > value {
		H.swap(i/2, i)
		i /= 2
	}
}

func (H *MinHeap) Print() {
	fmt.Println(H.data[1:])
}

//删除总是 将第一个(最值和最后一个元素交换,然后删除最后一个元素,再将1号元素下沉到正确位置)
func (H *MinHeap) DelMin() int {
	if len(H.data) > 1 {
		min := H.data[1]
		//顶和最后一个交换
		H.swap(1, len(H.data)-1)
		//删除最后一个
		H.data = H.data[:len(H.data)-1]
		//下沉逻辑
		i := 1
		for 2*i <= len(H.data)-1 {
			//假设左边最小
			min := 2 * i
			//存在右边,且右比左还小
			if 2*i+1 <= len(H.data)-1 && H.data[min] > H.data[2*i+1] {
				min = 2*i + 1
			}
			if H.data[i] < H.data[min] {
				// 根最小，不用交换
				break
			}
			H.swap(i, min)
			i = min
		}
		return min
	} else {
		return H.data[0]
	}
}

func (H *MinHeap) GetSize() int {
	return len(H.data) - 1

}
