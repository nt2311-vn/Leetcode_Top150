package rotate_arr

func Rotate(nums []int, k int) {
	shiftSlice := make([]int, len(nums))

	for i := range nums {
		shiftSlice[(i+k)%len(nums)] = nums[i]
	}

	for i, v := range shiftSlice {
		nums[i] = v
	}
}
