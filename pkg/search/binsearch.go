/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-27 06:25:33
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-27 06:41:34
 */
package search

func BinSearch(array []int, k int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := (left + right) / 2
		if array[mid] == k {
			return mid
		} else if array[mid] > k {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1

}
