package mergesortedarray

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	expected := []int{1, 2, 2, 3, 5, 6}
	actual := Merge(nums1, m, nums2, n)

	if len(expected) != len(actual) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}
}
