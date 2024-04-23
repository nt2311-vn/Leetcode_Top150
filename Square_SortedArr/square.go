package squaresortedarr

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	i, j, index := 0, len(nums)-1, len(nums)-1

	for i <= j {
		iSquare := nums[i] * nums[i]
		jSquare := nums[j] * nums[j]

		if iSquare > jSquare {
			result[index] = iSquare
			i++
		} else {
			result[index] = jSquare
			j--
		}
		index--
	}

	return result
}
