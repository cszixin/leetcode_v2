/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-28 03:13:11
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-28 04:22:10
 */
package array

// 回文问题
// 判断是否为回文
func IsPalinDrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 返回s中以left,right为中心的最长回文串
func GetPalinDrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

//得到s中最长的回文串
func LongestPalindrome(s string) string {
	output := ""
	for i := 0; i < len(s); i++ {
		s1 := GetPalinDrome(s, i, i)
		s2 := GetPalinDrome(s, i, i+1)
		if len(s1) > len(output) {
			output = s1
		}
		if len(s2) > len(output) {
			output = s2
		}
	}
	return output
}

//得到s中所有的最长的回文串
func LongestPalindromeV2(s string) []string {
	//当前最长回文串
	output := ""
	var res []string
	for i := 0; i < len(s); i++ {
		s1 := GetPalinDrome(s, i, i)
		s2 := GetPalinDrome(s, i, i+1)
		flag := false
		if len(s1) > len(output) {
			output = s1
			flag = true
		}
		if len(s2) > len(output) {
			output = s2
			flag = true
		}
		if flag {
			//找到了更长的回文串
			res = make([]string, 0)
			res = append(res, output)
		} else {
			tmp := make([]string, 0)
			if len(s1) > len(s2) {
				tmp = append(tmp, s1)
			} else if len(s1) < len(s2) {
				tmp = append(tmp, s2)
			} else {
				tmp = append(tmp, s1, s1)
			}
			// 如果当前找到的长度和目前已知回文串长度一样,则加入结果容器
			if len(res[0]) == len(tmp[0]) {
				res = append(res, tmp...)
			}
		}
	}
	return res
}
