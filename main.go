/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-26 06:10:50
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-26 06:15:09
 */
package main

import (
	"fmt"

	"leetcode/pkg/sort"
)

func main() {
	array := []int{1, 0, 0, 1}
	sort.QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
