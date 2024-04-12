package subarraysum

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minLen := len(nums) + 1
	sum := 0
	start := 0

	for end := 0; end < len(nums); end++ {
		sum += nums[end]

		for sum >= target {
			if minLen > end-start+1 {
				minLen = end - start + 1
			}

			sum -= nums[start]
			start++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}
