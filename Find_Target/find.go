package findtarget

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	result := []int{}

	lower, upper := findBound(nums, target)
	for i := lower; i <= upper; i++ {
		if nums[i] == target {
			result = append(result, i)
		}
	}

	return result
}

func findBound(nums []int, target int) (int, int) {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := int((i + j) / 2)

		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			break
		}
	}

	return i, j
}
