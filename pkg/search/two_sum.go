/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-27 07:02:13
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-27 07:06:38
 */
package search

func TwoSum(array []int, k int) []int {
	left, right := 0, len(array)-1
	for left <= right {
		value := array[left] + array[right]
		if value == k {
			return []int{left + 1, right + 1}
		} else if value > k {
			right--
		} else {
			left++
		}
	}
	return []int{-1, -1}
}
