package baseballgame

import "testing"

func TestOpsOne(t *testing.T) {
	ops := []string{"5", "2", "C", "D", "+"}
	expected := 30

	actual := callPoints(ops)
	if actual != expected {
		t.Errorf("Failed 1st test: expected %d but got %d", expected, actual)
	}
}

func TestOpsTwo(t *testing.T) {
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	expected := 27

	actual := callPoints(ops)
	if actual != expected {
		t.Errorf("Failed 2nd test: expected %d but got %d", expected, actual)
	}
}

func TestOpsThree(t *testing.T) {
	ops := []string{"1", "C"}
	expected := 0

	actual := callPoints(ops)
	if actual != expected {
		t.Errorf("Failed 3rd test: expected %d but got %d", expected, actual)
	}
}
