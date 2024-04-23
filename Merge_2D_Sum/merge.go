package merge2dsum

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	result := [][]int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] == nums2[j][0] {
			id := nums1[i][0]
			sum := nums1[i][1] + nums2[j][1]
			result = append(result, []int{id, sum})
			i, j = i+1, j+1
		} else if nums1[i][0] > nums2[j][0] {
			result = append(result, nums2[j])
			j++
		} else {
			result = append(result, nums1[i])
			i++
		}
	}

	if i < len(nums1) {
		result = append(result, nums1[i:]...)
	}

	if j < len(nums2) {
		result = append(result, nums2[j:]...)
	}

	return result
}
