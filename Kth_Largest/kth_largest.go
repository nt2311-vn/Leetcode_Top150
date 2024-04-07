package kthlargest

func partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	i := lo

	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[hi] = nums[hi], nums[i]

	return i
}

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	lo, hi := 0, len(nums)-1

	for lo < hi {
		index := partition(nums, lo, hi)

		if index < k {
			lo = index + 1
		} else if index > k {
			hi = index - 1
		} else {
			break
		}
	}

	return nums[k]
}
