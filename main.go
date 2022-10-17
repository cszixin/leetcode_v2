/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-26 06:10:50
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-14 16:46:32
 */
package main

import (
	"leetcode/pkg/queue"
)

func main() {
	q := queue.MakeQueue()
	q.Push(12)
	print(q.Max(), "\n")
	q.Push(9)
	print(q.Max(), "\n")
	q.Pop(12)
	q.Push(11)
	print(q.Max(), "\n")
	q.Push(8)
	print(q.Max(), "\n")
	q.Push(14)
	print(q.Max(), "\n")
	q.Pop(14)
	print(q.Max(), "\n")

}
