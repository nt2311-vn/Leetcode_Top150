package productexcludeself

func productExceptSelf(nums []int) []int {
	mem := len(nums)
	result := make([]int, mem)

	leftPointer := 1
	for i := 0; i < mem; i++ {
		result[i] = leftPointer
		leftPointer *= nums[i]
	}

	rightPointer := 1
	for i := mem - 1; i >= 0; i-- {
		result[i] *= rightPointer
		rightPointer *= nums[i]
	}

	return result
}
