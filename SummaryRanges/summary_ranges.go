package summaryranges

import "fmt"

func summaryRanges(nums []int) []string {
	result := []string{}

	if len(nums) == 0 {
		return result
	}

	for i := 0; i < len(nums); i++ {
		start := nums[i]

		for i+1 < len(nums) && nums[i+1] == nums[i]+1 {
			i++
		}

		if start == nums[i] {
			result = append(result, fmt.Sprintf("%d", start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, nums[i]))
		}
	}

	return result
}
