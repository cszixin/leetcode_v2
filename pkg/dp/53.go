package dp

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	//初始化
	copy(dp, nums)
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		//迭代dp
		dp[i] = max(dp[i-1]+dp[i], dp[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
