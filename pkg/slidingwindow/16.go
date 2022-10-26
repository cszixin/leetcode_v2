/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-17 14:18:53
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-17 14:22:05
 */
package slidingwindow

import "math"

func lengthOfLongestSubstringv2(s string) int {
	windows := make(map[byte]int)
	left, right := 0, 0
	maxLen := math.MinInt
	for right < len(s) {
		c := s[right]
		right++
		windows[c]++
		for windows[c] > 1 {
			d := s[left]
			left++
			windows[d]--
		}
		if right-left > maxLen {
			maxLen = right - left
		}

	}

	if maxLen == math.MinInt {
		return 0
	}
	return maxLen

}
