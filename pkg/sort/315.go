package sort

type pair struct {
	// 数组中原索引
	key int
	// 数组元素的值
	value int
}

func countSmaller(nums []int) []int {
	res := make([]int, len(nums))
	// 转化为k,v对
	arr := make([]pair, len(nums))
	for k, v := range nums {
		arr[k] = pair{k, v}
	}

	tmp := make([]pair, len(nums))

	merge := func(arr []pair, start, mid, end int) {

		for p := start; p <= end; p++ {
			tmp[p] = arr[p]
		}

		i, j := start, mid+1
		for p := start; p <= end; p++ {
			if i == mid+1 {
				arr[p] = tmp[j]
				j++
			} else if j == end+1 {
				arr[p] = tmp[i]
				res[arr[p].key] += j - (mid + 1)
				i++
			} else if tmp[i].value <= tmp[j].value {
				arr[p] = tmp[i]
				// 此时[mid+1,j-1]是符合要求的
				res[arr[p].key] += j - (mid + 1)
				i++
			} else {
				arr[p] = tmp[j]
				j++
			}

		}

	}
	var mysort func([]pair, int, int)
	mysort = func(nums []pair, start, end int) {
		if start >= end {
			return
		}
		mid := start + (end-start)/2
		mysort(nums, start, mid)
		mysort(nums, mid+1, end)
		merge(nums, start, mid, end)
	}

	mysort(arr, 0, len(arr)-1)
	return res
}
