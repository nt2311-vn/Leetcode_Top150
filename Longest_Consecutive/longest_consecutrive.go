package longestconsecutive

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	hashMap := make(map[int]bool, len(nums))
	for _, val := range nums {
		hashMap[val] = true
	}

	countConsecutive := 0

	for _, num := range nums {

		if hashMap[num-1] {
			continue
		}

		sequence, nextSequenxe := 1, num+1

		for hashMap[nextSequenxe] {
			sequence++
			nextSequenxe++
		}

		if sequence > countConsecutive {
			countConsecutive = sequence
		}
	}

	return countConsecutive
}
