package goodpairs

func numIdenticalPairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	count := 0

	for i := 0; i < len(nums); i++ {
		j := i + 1

		for j < len(nums) {
			if nums[j] == nums[i] {
				count++
			}
			j++
		}
	}
	return count
}
