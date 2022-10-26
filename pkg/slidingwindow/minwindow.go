/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-13 14:16:42
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-17 10:06:41
 */
package slidingwindow

import (
	"leetcode/pkg/queue"
	"math"
)

func minWindow(s string, t string) string {
	// 滑动window的左右边界 [left, right)，即不包含右边
	left, right := 0, 0
	valid := 0
	minLen := math.MaxInt
	start := -1
	needs := make(map[byte]int)
	windows := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		needs[t[i]] = needs[t[i]] + 1
	}

	for right < len(s) {
		c := s[right]
		// 扩大窗口
		right++
		// 调整窗口内的数据,是我关心的数据
		if _, ok := needs[c]; ok {
			//进入窗口
			windows[c] = windows[c] + 1
			if needs[c] == windows[c] {
				valid++
			}
		}
		//找到一个满足条件的
		for valid == len(needs) {
			// 缩小窗口
			//更新最小子串
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			d := s[left]
			left++
			if _, ok := needs[d]; ok {
				if windows[d] == needs[d] {
					//退出窗口了
					valid--
				}
				windows[d] = windows[d] - 1
			}
		}
	}
	if start == -1 {
		return ""
	}
	return s[start : start+minLen]
}

func checkInclusion(t, s string) bool {
	left, right := 0, 0
	valid := 0
	needs, windows := make(map[byte]int, len(t)), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		needs[t[i]] = needs[t[i]] + 1
	}
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := needs[c]; ok {
			windows[c] = windows[c] + 1
			if windows[c] == needs[c] {
				valid++
			}
		}
		//当滑动窗口的长度==t的长度是，开始缩小
		if right-left == len(t) {
			// 将全排问题转化为，长度相同，而且每个字符的个数也相同
			if valid == len(needs) {
				return true
			}
			d := s[left]
			left++
			if _, ok := needs[d]; ok {
				if windows[d] == needs[d] {
					valid--
				}
				windows[d] = windows[d] - 1
			}

		}

	}
	return false
}

// 在s中找p的异位词
func findAnagrams(s string, p string) []int {
	left, right, valid := 0, 0, 0
	res := make([]int, 0)
	needs, windows := make(map[byte]int, len(p)), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		needs[p[i]] = needs[p[i]] + 1
	}
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := needs[c]; ok {
			windows[c]++
			if needs[c] == windows[c] {
				valid++
			}
		}

		for right-left == len(p) {
			if valid == len(needs) {
				// 找到
				res = append(res, left)
			}
			// 缩减窗口
			d := s[left]
			left++
			if _, ok := needs[d]; ok {
				if windows[d] == needs[d] {
					valid--
				}
				windows[d] = windows[d] - 1
			}
		}
	}
	return res
}

// 最长无重复子串的长度
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	res := 0
	windows := make(map[byte]int)
	for right < len(s) {
		c := s[right]
		//加入窗口
		right++
		// 窗口内数据更新
		windows[c] = windows[c] + 1
		//是否要缩小窗口
		for windows[c] > 1 {
			d := s[left]
			left++
			windows[d]--
		}
		//更新结果
		if right-left > res {
			res = right - left
		}
	}
	return res
}

// leetcode 1004题, 将题意 转换一个，即子数组里面不超过k个0的 最长子串
func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	max := math.MinInt
	windows := make(map[int]int)
	for right < len(nums) {
		c := nums[right]
		right++
		windows[c]++

		for windows[0] > k {
			d := nums[left]
			windows[d]--
			left++
		}
		// 更新值,此时肯定是连续的且不超过k个0
		if windows[0]+windows[1] > max {
			max = windows[0] + windows[1]
		}
	}
	return max
}

// TODO 用单调队列优化
func getMaxDiff(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max, min := nums[0], nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	diff := max - min
	return diff
}

func longestSubarray(nums []int, limit int) int {
	left, right := 0, 0
	max := 0
	q1 := queue.MakeQueue(true)
	q2 := queue.MakeQueue(false)

	for right < len(nums) {
		c := nums[right]
		right++
		q1.Push(c)
		q2.Push(c)

		for q1.Max()-q2.Max() > limit {
			q1.Pop(nums[left])
			q2.Pop(nums[left])
			left++
		}
		if right-left > max {
			max = right - left
		}
	}
	return max
}

// leetcode 1658 转化问题：求连续子数组的和为sum-x 且长度最长
func minOperations(nums []int, x int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	target := sum - x
	if target < 0 {
		return -1
	}
	sum = 0
	left, right := 0, 0
	maxLen := 0
	flag := false
	for right < len(nums) {
		c := nums[right]
		right++
		sum += c
		for sum > target {
			d := nums[left]
			left++
			sum -= d
		}
		if sum == target {
			flag = true
			if right-left > maxLen {
				maxLen = right - left
			}
		}
	}
	if !flag {
		return -1
	}
	return len(nums) - maxLen
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		//进入单调队列,比进入元素小的全部出去
		for j := len(queue) - 1; j >= 0; j-- {
			if nums[i] > queue[j] {
				queue = queue[:len(queue)-1]
			} else {
				break
			}
		}
		// for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
		// 	queue = queue[:len(queue)-1]
		// }
		queue = append(queue, nums[i])
		//窗口大小达到k
		if i >= k-1 {
			//当前窗口的最大值，即为队头
			res = append(res, queue[0])
			if queue[0] == nums[i+1-k] {
				queue = queue[1:]
			}
		}
	}
	return res
}
