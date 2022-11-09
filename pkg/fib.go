package pkg

func fibv1(N int) int {
	if N == 1 || N == 2 {
		return N
	}
	return fibv1(N-1) + fibv1(N-2)
}

func fibv2(N int) int {
	memo := make([]int, N+1)
	var help func(N int) int
	help = func(N int) int {
		if N == 0 || N == 1 {
			return N
		}
		//在备忘录中查询结果
		if memo[N] != 0 {
			return memo[N]
		}
		//记录结果
		memo[N] = help(memo[N-1]) + help(memo[N-2])
		return memo[N]
	}
	return help(N)
}

func fibv3(N int) int {
	if N == 0 {
		return N
	}

	dp := make([]int, N+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]

}

func fibv4(N int) int {
	if N == 0 {
		return N
	}
	dp_i_1 := 1
	dp_i_2 := 0

	for i := 2; i <= N; i++ {
		dp_i := dp_i_1 + dp_i_2
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i_1
}
