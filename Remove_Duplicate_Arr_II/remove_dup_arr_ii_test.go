package removeduplicatearrii

import "testing"

func TestRemoveDupOne(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	const expected = 5

	actual := RemoveDuplicates(nums)

	if actual != expected {
		t.Errorf("Expected as %d but actual %d\n", expected, actual)
	}
}

func TestRemoveDupTwo(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	const expected = 7

	actual := RemoveDuplicates(nums)
	if actual != expected {
		t.Errorf("Expected as %d but actual %d\n", expected, actual)
	}
}
