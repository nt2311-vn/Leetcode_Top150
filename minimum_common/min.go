package minimumcommon

import "math"

func getCommon(nums1 []int, nums2 []int) int {
	common := []int{}

	for _, num := range nums1 {
		if binarySearch(nums2, num) != -1 {
			common = append(common, num)
		}
	}

	if len(common) == 0 {
		return -1
	} else {
		return findMin(common)
	}
}

func binarySearch(nums []int, target int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		mid := i + (j-i)/2

		if nums[mid] > target {
			j = mid - 1
		} else if nums[mid] < target {
			i = mid + 1
		} else {
			return nums[mid]
		}
	}
	return -1
}

func findMin(nums []int) int {
	currentMin := math.MaxInt64

	for _, num := range nums {
		if num < currentMin {
			currentMin = num
		}
	}

	return currentMin
}
