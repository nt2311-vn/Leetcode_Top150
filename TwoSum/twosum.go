package twosum

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	result := []int{}
	for index, val := range nums {
		remain := target - val
		hashIndex, exist := hashMap[remain]

		if exist {
			return []int{hashIndex, index}
		} else {
			hashMap[val] = index
		}
	}

	return result
}
