package arrayconcatenation

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	result := 0
	for len(nums) > 0 {
		i, j := 0, len(nums)-1

		if i != j {
			concat := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])
			concatVal, _ := strconv.Atoi(concat)
			result += concatVal
		} else {
			result += nums[i]
		}

		if len(nums) > 1 {
			nums = nums[1 : len(nums)-1]
		} else {
			break
		}
	}

	return int64(result)
}
