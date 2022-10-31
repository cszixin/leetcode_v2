package sort

func CallMergeSort(nums []int) []int {
	MergerSort(nums, 0, len(nums)-1)
	return nums
}

func MergerSort(nums []int, start, end int) {
	//递归出口
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	//左边
	MergerSort(nums, start, mid)
	MergerSort(nums, mid+1, end)
	merge(nums, start, mid, end)
}

func merge(nums []int, start, mid, end int) {
	tmp := make([]int, end-start+1)
	// 左边起点
	i := start
	//右边起点
	j := mid + 1
	k := 0

	for i <= mid && j <= end {
		if nums[i] > nums[j] {
			tmp[k] = nums[j]
			j++
		} else {
			tmp[k] = nums[i]
			i++
		}
		k++
	}
	//左边已经遍历完,复制右边剩下的
	for j <= end {
		tmp[k] = nums[j]
		j++
		k++
	}
	//右边已经遍历完,复制左边边剩下的
	for i <= mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	// 复制tmp到nums [start,end]中
	for k := 0; k < len(tmp); k++ {
		nums[start+k] = tmp[k]
	}
}
