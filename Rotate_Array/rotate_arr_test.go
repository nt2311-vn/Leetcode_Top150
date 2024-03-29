package rotate_arr

import "testing"

func TestRotateOne(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	expected := []int{5, 6, 7, 1, 2, 3, 4}
	Rotate(nums, k)

	if len(nums) != len(expected) {
		t.Errorf("Expected slice len %d but got %d\n", len(expected), len(nums))
	}

	for i := range expected {
		if expected[i] != nums[i] {
			t.Errorf("Expected %v but got %v\n", expected, nums)
		}
	}
}
