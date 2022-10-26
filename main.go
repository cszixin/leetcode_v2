/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-26 06:10:50
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-19 16:56:11
 */
package main

import (
	"container/heap"
	"fmt"
	myheap "leetcode/pkg/heap"
	"sync"
)

func doIter(m sync.Map) {
	m.Range(func(key, value interface{}) bool {
		k, _ := key.(int)
		if v, ok := value.(int); ok && k%2 == 0 {
			m.Store(key, v+1)

		}
		return true
	})

}

func doIterV2(m sync.Map) {
	m.Range(func(key, value interface{}) bool {
		k, _ := key.(int)
		if v, ok := value.(int); ok && k%2 == 1 {
			m.Store(key, v+1)

		}
		return true
	})

}

// func doIterv2(m sync.Map) {
// 	for k, v := range m {
// 		_ = fmt.Sprintf("%d, %d", k, v)

// 	}
// }

// func dowrite(m sync.Map) {
// 	for k, v := range m {
// 		m[k] = v + 1
// 	}
// }
func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5 // 1.x = 5 2. defer x = 6 3  真正的返回
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中x的副本
	}(x)
	return 5 // 返回值 = x = 5
}

func main() {
	h := make(myheap.MaxHeap, 0)
	heap.Init(&h)
	heap.Push(&h, 9)
	// heap.Push(&h, 1)
	// heap.Push(&h, 4)
	// heap.Push(&h, 10)
	// heap.Push(&h, 11)
	fmt.Println(heap.Pop(&h))
}
