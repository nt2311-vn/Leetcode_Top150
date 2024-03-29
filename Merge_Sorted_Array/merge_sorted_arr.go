package mergesortedarray

import "sort"

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	nums1 = append(nums1[0:m], nums2[0:n]...)

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] <= nums1[j]
	})

	return nums1
}
