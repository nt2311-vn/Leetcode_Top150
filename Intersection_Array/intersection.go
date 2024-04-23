package intersectionarray

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	hashMap := map[int]bool{}

	sort.Ints(nums1)
	sort.Ints(nums2)

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			if !hashMap[nums1[i]] {
				result = append(result, nums1[i])
				hashMap[nums1[i]] = true
			}
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}

	return result
}
