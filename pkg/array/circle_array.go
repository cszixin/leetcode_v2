package array

func PrintCircleArray(array []int) []int {
	output := make([]int, 0)
	orgLen := len(array)
	for i := 0; i <= 2*orgLen-1; i++ {
		output = append(output, array[i%orgLen])
	}
	return output
}
