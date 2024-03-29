package major_element

import "testing"

func TestMajorOne(t *testing.T) {
	nums := []int{3, 2, 3}
	expected := 3

	actual := MajorityElement(nums)

	if actual != expected {
		t.Errorf("Failed 1st test: expected %d but got %d\n", expected, actual)
	}
}

func TestMajorTwo(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	expected := 2

	actual := MajorityElement(nums)

	if actual != expected {
		t.Errorf("Failed 2nd test: expeted %d but got %d", expected, actual)
	}
}
