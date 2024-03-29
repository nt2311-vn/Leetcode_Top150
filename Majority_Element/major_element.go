package major_element

func MajorityElement(nums []int) int {
	majorPoint := len(nums) / 2
	count := map[int]int{}

	for _, v := range nums {
		count[v]++

		if count[v] > majorPoint {
			return v
		}
	}

	panic("You promise it always has a major element")
}
