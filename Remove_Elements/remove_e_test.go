package removeelements

import "testing"

func TestRemoveOne(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3

	expected := 2
	actual := RemoveElement(nums, val)

	if expected != actual {
		t.Errorf("Expected %d but %d", expected, actual)
	}
}

func TestRemoveTwo(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	expected := 5
	actual := RemoveElement(nums, val)

	if expected != actual {
		t.Errorf("Expected %d but %d", expected, actual)
	}
}
