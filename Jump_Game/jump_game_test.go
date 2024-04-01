package jump_game

import "testing"

func TestFirstCase(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	expected := true

	actual := canJump(nums)
	if actual != expected {
		t.Errorf("Failed 1st test: expected %v but got %v\n", expected, actual)
	}
}

func TestSecondCase(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	expected := false

	actual := canJump(nums)
	if actual != expected {
		t.Errorf("Failed 2nd test: expected %v but got %v\n", expected, actual)
	}
}

func TestThirdCase(t *testing.T) {
	nums := []int{2, 0}
	expected := true

	actual := canJump(nums)
	if actual != expected {
		t.Errorf("Failed 3rd test: expected %v but got %v\n", expected, actual)
	}
}

func TestFourthCase(t *testing.T) {
	nums := []int{2, 0, 0}
	expected := true

	actual := canJump(nums)
	if actual != expected {
		t.Errorf("Failed 4th test: expected %v but got %v\n", expected, actual)
	}
}

func TestFifThCase(t *testing.T) {
	nums := []int{1, 1, 1, 0}
	expected := true

	actual := canJump(nums)
	if actual != expected {
		t.Errorf("Failed 5th test: expected %v but got %v\n", expected, actual)
	}
}
