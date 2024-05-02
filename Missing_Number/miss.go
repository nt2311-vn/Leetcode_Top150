package missingnumber

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)

	for i < j {
		mid := i + (j-i)/2

		if nums[mid] > mid {
			j = mid
		} else {
			i = mid + 1
		}
	}

	return i
}
