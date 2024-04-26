package maxcountpostivenegative

func maximumCount(nums []int) int {
	negative, positve := findBound(nums)
	result := getMax(negative, positve)

	return result
}

func findBound(nums []int) (int, int) {
	i, j := 0, len(nums)-1

	for i <= j {
		mid := int(i + (j-i)/2)
		if nums[mid] < 0 {
			i = mid + 1
		} else if nums[mid] > 0 {
			j = mid - 1
		} else {
			i = mid - 1
			j = mid + 1
		}
	}

	return i, len(nums) - j - 1
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
