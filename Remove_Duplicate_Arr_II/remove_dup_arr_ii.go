package removeduplicatearrii

func RemoveDuplicates(nums []int) int {
	count_Dup := map[int]int{}

	for i := len(nums) - 1; i >= 0; i-- {
		count_Dup[nums[i]]++

		if count_Dup[nums[i]] >= 3 {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}

	return len(nums)
}
