package dp

var mymemo [][]int

func longestCommonSubsequence(text1 string, text2 string) int {
	for row := 0; row < len(text1); row++ {
		tmp := make([]int, 0)
		for col := 0; col < len(text2); col++ {
			tmp = append(tmp, -1)
		}
		mymemo = append(mymemo, tmp)
	}
	res := lcs(text1, 0, text2, 0)
	mymemo = make([][]int, 0)
	return res

}

// s1[i...]和s2[j...]的LCS
func lcs(s1 string, i int, s2 string, j int) int {

	// base case
	if i == len(s1) || j == len(s2) {
		return 0
	}
	if mymemo[i][j] != -1 {
		return mymemo[i][j]
	}

	if s1[i] == s2[j] {
		mymemo[i][j] = 1 + lcs(s1, i+1, s2, j+1)

	} else {
		mymemo[i][j] = max(lcs(s1, i, s2, j+1), lcs(s1, i+1, s2, j))
	}
	return mymemo[i][j]
}
