package h_index

import "testing"

func TestFirstCase(t *testing.T) {
	citations := []int{3, 0, 6, 1, 5}
	expected := 3

	actual := hIndex(citations)
	if expected != actual {
		t.Errorf("Failed 1st test: expected %d but got %d\n", expected, actual)
	}
}

func TestSecondCase(t *testing.T) {
	citations := []int{1, 3, 1}
	expected := 1

	actual := hIndex(citations)
	if expected != actual {
		t.Errorf("Failed 2nd test: expected %d but got %d\n", expected, actual)
	}
}
