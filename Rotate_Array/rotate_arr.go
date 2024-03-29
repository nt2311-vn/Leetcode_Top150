package rotate_arr

func Rotate(nums []int, k int) {
	for i := 1; i <= k; i++ {
		nums[i-1] = nums[len(nums)-i]
	}
}
