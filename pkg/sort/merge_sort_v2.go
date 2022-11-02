package sort

func MergerSortV2(nums []int) {
	tmp := make([]int, len(nums))
	var mysort func([]int, int, int)

	merge := func(nums []int, start, mid, end int) {
		// 先整体移入tmp数组
		for i := start; i <= end; i++ {
			tmp[i] = nums[i]
		}
		left := start
		right := mid + 1
		// 填充nums[start,end],左右两边都在tmp数组了
		for p := start; p <= end; p++ {
			if left == mid+1 {
				//左边已经完成
				nums[p] = tmp[right]
				right++
			} else if right == end+1 {
				//右边已经完成
				nums[p] = tmp[left]
				left++
			} else if tmp[left] > tmp[right] {
				nums[p] = tmp[right]
				right++
			} else {
				nums[p] = tmp[left]
				left++
			}

		}

	}
	mysort = func(nums []int, start, end int) {
		if start >= end {
			return
		}
		mid := start + (end-start)/2
		mysort(nums, start, mid)
		mysort(nums, mid+1, end)
		merge(nums, start, mid, end)
	}
	mysort(nums, 0, len(nums)-1)
}
