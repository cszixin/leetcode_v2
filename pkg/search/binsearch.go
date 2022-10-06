/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-27 06:25:33
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-29 16:26:10
 */
package search

func BinSearch(array []int, k int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)/2
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

func BinSearchLeftBound(array []int, target int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] > target {
			right = mid - 1
		} else if array[mid] < target {
			left = mid + 1
		} else if array[mid] == target {
			//找到了一个
			if mid == 0 || array[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func BinSearchRightBound(array []int, target int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)/2
		if array[mid] < target {
			left = mid + 1
		} else if array[mid] > target {
			right = mid - 1
		} else if array[mid] == target {
			if mid == len(array)-1 || array[mid+1] != target {
				return mid
			} else {
				left = mid + 1
			}

		}
	}
	return -1
}

func BinSearch33(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] >= nums[0] {
			// 左边较大部分
			if nums[0] <= target && nums[mid] > target {
				//target在左边区域
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			//右边较小部分
			if nums[mid] < target && nums[len(nums)-1] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			//找到一个比target小的
			// 已经到达最后，或者mid后面一个比target大,说明mid已经是最后一个比target小的
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid + 1
			} else {
				left = mid + 1
			}
		}
	}
	// left=right+1,到达了最左端
	return 0
}

func MySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			// 大了
			right = mid - 1
		} else if mid*mid < x {
			// 找到一个小的
			if mid == x || (mid+1)*(mid+1) > x {
				return mid
			} else {
				left = mid + 1
			}
		}
	}
	return 0
}
