package containduplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	indexTable := map[int]int{}

	for index, num := range nums {
		tableIndex, exist := indexTable[num]

		if exist {
			if index-tableIndex <= k {
				return true
			}
		}
		indexTable[num] = index
	}

	return false
}
