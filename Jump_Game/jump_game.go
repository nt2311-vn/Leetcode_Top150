package jump_game

func canJump(nums []int) bool {
	maxSteps := 0
	for i := 0; i < len(nums); i++ {
		if i > maxSteps {
			return false
		}

		if i+nums[i] > maxSteps {
			maxSteps = i + nums[i]
		}
	}

	return true
}
