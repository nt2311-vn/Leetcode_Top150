package major_element

// func MajorityElement(nums []int) int {
// 	majorPoint := len(nums) / 2
// 	count := map[int]int{}
//
// 	for _, v := range nums {
// 		count[v]++
//
// 		if count[v] > majorPoint {
// 			return v
// 		}
// 	}
//
// 	panic("You promise it always has a major element")
// }

func MajorityElement(nums []int) int {
	result, count := nums[0], 1

	for _, v := range nums[1:] {
		if result == v {
			count++
		} else {
			count--
		}

		if count < 0 {
			result = v
			count = 0
		}

	}

	return result
}
