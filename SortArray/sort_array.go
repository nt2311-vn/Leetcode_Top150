package sortarray

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot := nums[len(nums)/2]

	var left, right []int

	for _, num := range nums {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	return append(append(sortArray(left), pivot), sortArray(right)...)
}
