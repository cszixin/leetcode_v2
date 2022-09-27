/*
 * @Descripttion:
 * @version:
 * @Author: liuchuanshi
 * @Date: 2022-09-27 03:40:01
 * @LastEditors: liuchuanshi
 * @LastEditTime: 2022-09-27 04:20:08
 */
package array

func PrintCircleArray(array []int) []int {
	output := make([]int, 0)
	orgLen := len(array)
	for i := 0; i <= 2*orgLen-1; i++ {
		output = append(output, array[i%orgLen])
	}
	return output
}

func RemoveDuplicates(array []int) int {
	if len(array) <= 1 {
		return len(array)
	}
	slow, fast := 0, 0
	for fast < len(array) {
		// 找到一个合适的成员
		if array[fast] != array[slow] {
			slow++
			array[slow] = array[fast]
		}
		fast++
	}
	// 新数组的长度
	return slow + 1
}

func RemoveElement(array []int, val int) int {
	fast, slow := 0, 0
	for fast < len(array) {
		if array[fast] != val {
			array[slow] = array[fast]
			slow++
		}
		fast++
	}
	return slow
}
