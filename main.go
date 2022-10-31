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
	"fmt"
	"sort"
)

func main() {
	type pair struct{ score, idx int }

	a := []pair{{10, 0}, {9, 1}, {12, 2}, {11, 3}, {12, 4}}
	sort.SliceStable(a, func(i, j int) bool { return a[i].score < a[j].score })
	fmt.Println(a)
}
