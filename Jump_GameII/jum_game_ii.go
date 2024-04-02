package jumpgameii

func jump(nums []int) int {
	steps, end, index := 0, 0, 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > index {
			index = i + nums[i]
		}

		if i == end {
			end = index
			steps++
		}
	}
	return steps
}
