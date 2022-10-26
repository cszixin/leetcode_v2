/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-10-19 09:38:48
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-10-19 11:29:31
 */
package sort

import "math/rand"

func QuickSort(array []int, low, high int) {
	if low >= high {
		return
	}
	provit := partition(array, low, high)
	QuickSort(array, 0, provit-1)
	QuickSort(array, provit+1, high)
}

func swap(array []int, a, b int) {
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}

func swapv2(array []int, a, b int) {
	// temp := array[a]
	// array[a] = array[b]
	// array[b] = temp
	array[a], array[b] = array[b], array[a]
}

func partition(array []int, low, high int) int {
	// 基准
	start := low
	p := array[low]
	for low < high {
		for high > low && array[high] > p {
			high--
		}
		for high > low && array[low] <= p {
			low++
		}
		swap(array, low, high)
	}
	// low与high重合后，基准的位置确定
	swap(array, start, low)
	return high
}

func partitionv2(nums []int, low, high int) int {

	randNum := rand.Intn(high-low) + low
	swapv2(nums, low, randNum)
	p := nums[low]
	start := low
	for low < high {
		for low < high && nums[high] > p {
			high--
		}
		for low < high && nums[low] <= p {
			low++
		}
		swapv2(nums, low, high)

	}
	swapv2(nums, start, low)
	return low
}

func QuickSortV2(nums []int) []int {

	var sort func(nums []int, low, high int)
	sort = func(nums []int, low, high int) {

		if low >= high {
			return
		}
		p := partitionv2(nums, low, high)
		sort(nums, low, p-1)
		sort(nums, p+1, high)

	}
	sort(nums, 0, len(nums)-1)
	return nums
}

func findKthLargest(nums []int, k int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		p := partition(nums, low, high)
		if p == len(nums)-k {
			//倒数第k的元素，就是第k大的
			return nums[p]
		} else if p > len(nums)-k {
			//第k大的在左边
			high = p - 1
		} else {
			//在右边
			low = p + 1
		}
	}
	return -1
}
