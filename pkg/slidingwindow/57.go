/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-17 10:24:29
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-17 10:32:12
 */
package slidingwindow

func findContinuousSequence(target int) [][]int {
	left, right := 0, 0
	sum := 0
	res := make([][]int, 0)
	for right < target {
		c := right + 1
		right++
		sum += c
		for sum > target {
			d := left + 1
			left++
			sum -= d
		}
		if sum == target && right-left > 1 {
			tmp := make([]int, 0)
			for k := left; k < right; k++ {
				tmp = append(tmp, k+1)
			}
			res = append(res, tmp)
		}
	}
	return res
}
