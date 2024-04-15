package buildsarr

func buildArray(nums []int) []int {
	result := make([]int, len(nums))

	if len(nums) == 0 {
		return result
	}

	for i := 0; i < len(nums); i++ {
		result[i] = nums[nums[i]]
	}

	return result
}
