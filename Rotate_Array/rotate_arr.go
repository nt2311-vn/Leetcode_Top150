package rotate_arr

func Rotate(nums []int, k int) {
	if k < len(nums) {
		copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	} else {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
