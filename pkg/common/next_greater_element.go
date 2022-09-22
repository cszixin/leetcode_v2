package common

func NextGreaterElement(array []int) []int {
	ans := make([]int, len(array))
	stack := make([]int, len(array))
	for i := len(array) - 1; i >= 0; i-- {
		// stack不空，当前迭代元素比栈顶大,则栈顶出栈
		for len(stack) > 0 && array[i] >= stack[len(stack)-1] {
			//出栈
			stack = stack[:len(stack)-1]
		}
		// 跳出循环,说明后面没有更大的元素
		if len(stack) == 0 {
			ans[i] = -1
		} else {
			//比他大且最近的元素，就是是栈顶
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, array[i])
	}
	return ans
}

// 应用 https://leetcode.cn/problems/daily-temperatures/
func DailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			// 出栈
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = 0
		} else {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return ans

}
