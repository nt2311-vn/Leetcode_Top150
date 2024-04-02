package productexcludeself

import "testing"

func TestFirstCase(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{24, 12, 8, 6}

	actual := productExceptSelf(nums)

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("Failed first test: expected %v but got %v\n", expected, actual)
		}
	}
}

func TestSecondCase(t *testing.T) {
	nums := []int{-1, 1, 0, -3, 3}
	expected := []int{0, 0, 9, 0, 0}

	actual := productExceptSelf(nums)

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("Failed second test: expected %v but got %v\n", expected, actual)
		}
	}
}
