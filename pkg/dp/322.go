package dp

import "math"

var memo []int

func coinChange(coins []int, amount int) int {
	memo = make([]int, amount+1)
	memo[0] = 0
	return dp(coins, amount)

}

func dp(coins []int, amount int) int {

	// base case
	if amount < 0 {
		// 不合法
		return -1
	} else if amount == 0 {
		return memo[0]
	}
	if memo[amount] != 0 {
		return memo[amount]
	}
	res := math.MaxInt
	for _, coin := range coins {
		sub := dp(coins, amount-coin)
		if sub == -1 {
			// 子问题无解
			continue
		}
		res = min(res, sub+1)
	}

	// 要在此处更新memo数组
	if res == math.MaxInt {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a

}

func coinChangev2(coins []int, amount int) int {
	// dp[i] 表示 i元钱 至少需要几个硬币来凑
	dp := make([]int, amount+1)
	for k := range dp {
		dp[k] = amount + 1
	}
	// base case 0元当然需要0个硬币
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		// 对于
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}

	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]

}
