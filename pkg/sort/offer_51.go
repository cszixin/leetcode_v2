package sort

// offer_51和 493是同一类题型
func reversePairs(nums []int) int {
	// 创建辅助切片,避免后续频繁创建
	tmp := make([]int, len(nums))
	count := 0
	var mysort func([]int, int, int)
	// var merge func([]int, int, int, int)
	merge := func(nums []int, start, mid, end int) {

		for p := start; p <= end; p++ {
			// nums[start, end]整体复制到tmp
			tmp[p] = nums[p]
		}
		i := start
		j := mid + 1
		// 携带私货
		k := mid + 1 // 此句不可放入for内，会timeout, 当p 右移时,k不必回头
		for p := start; p <= mid; p++ {
			for k <= end && tmp[p] > 2*tmp[k] {
				k++
			}
			count += k - (mid + 1)
		}

		for p := start; p <= end; p++ {
			// 填充nums
			if i == mid+1 {
				// 左边已经扫描完毕
				nums[p] = tmp[j]
				j++
			} else if j == end+1 {
				//右边已经扫描完毕
				nums[p] = tmp[i]
				i++
			} else if tmp[i] <= tmp[j] {
				nums[p] = tmp[i]
				i++
			} else {
				nums[p] = tmp[j]
				j++
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
	return count
}
