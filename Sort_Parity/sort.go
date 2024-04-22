package sortparity

func sortArrayByParity(nums []int) []int {
	result := make([]int, len(nums))

	i, j := 0, len(nums)-1

	for _, num := range nums {
		if num%2 == 0 {
			result[i] = num
			i++
		} else {
			result[j] = num
			j--
		}
	}

	return result
}
