package sortcolors

func sortColors(nums []int) {
	bucket := make([]int, 3)

	for _, num := range nums {
		bucket[num] += 1
	}

	i := 0
	for j := range bucket {
		for bucket[j] > 0 {
			nums[i] = j
			bucket[j]--
			i++
		}
	}
}
