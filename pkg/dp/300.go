package dp

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {

	// base case
	if len(nums) == 0 {
		return 0
	}
	// dp[i] 以nums[i]结尾的的最长递增子序列的长度
	dp := make([]int, len(nums))
	//紧紧包含自己
	var i, j int
	dp[0] = 1
	for i = 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i = 1; i < len(nums); i++ {
		for j = i - 1; j >= 0; j-- {
			// 往回找
			if nums[j] < nums[i] {
				// 找到符合条件的
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := dp[0]
	for i = 0; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
