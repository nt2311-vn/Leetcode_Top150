package missingnumber

import "sort"

func missingNumber(nums []int) int {
	for i := len(nums); i >= 0; i-- {
		if !binarySearch(nums, i) {
			return i
		}
	}

	return 0
}

func binarySearch(nums []int, target int) bool {
	sort.Ints(nums)
	i, j := 0, len(nums)-1

	for i <= j {
		mid := i + (j-i)/2

		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			return true
		}
	}

	return false
}
