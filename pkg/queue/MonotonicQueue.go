/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-14 16:13:10
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-17 09:45:39
 */
package queue

import "math"

//实现单调队列
type MonotonicQueue struct {
	data []int
	flag bool
}

func MakeQueue(flag bool) *MonotonicQueue {
	// flag 为true单调递减,求最大值,反之单调递增,求最小值
	q := new(MonotonicQueue)
	q.data = make([]int, 0)
	q.flag = flag
	return q
}

func (q *MonotonicQueue) Push(elem int) {

	for _, v := range q.data {
		if (q.flag && elem > v) || (!q.flag && elem < v) {
			q.data = q.data[:len(q.data)-1]
		} else {
			break
		}
	}
	q.data = append(q.data, elem)
}

func (q *MonotonicQueue) Pop(elem int) {
	if len(q.data) > 0 && q.data[0] == elem {
		q.data = q.data[1:]
	}
}

func (q *MonotonicQueue) Max() int {
	if len(q.data) == 0 {
		return math.MinInt
	}
	return q.data[0]
}
