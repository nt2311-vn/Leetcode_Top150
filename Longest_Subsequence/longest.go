package longestsubsequence

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	result := make([]int, len(queries))

	for index, num := range queries {
		sum, subsequenceCount := 0, 0

		for j := 0; j < len(nums); j++ {
			if sum+nums[j] <= num {
				sum += nums[j]
				subsequenceCount++
			} else {
				break
			}
		}

		result[index] = subsequenceCount
	}

	return result
}
