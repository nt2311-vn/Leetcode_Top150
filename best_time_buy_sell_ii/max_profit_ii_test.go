package besttimebuysellii

import "testing"

func TestFirstCase(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 7

	actual := maxProfit(prices)
	if expected != actual {
		t.Errorf("Failed first test: expected %d but got %d\n", expected, actual)
	}
}

func TestSecondCase(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5}
	expected := 4

	actual := maxProfit(prices)
	if expected != actual {
		t.Errorf("Failed second test: expected %d but got %d\n", expected, actual)
	}
}

func TestThirdCase(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expected := 0

	actual := maxProfit(prices)
	if expected != actual {
		t.Errorf("Failed third test: expected %d but got %d\n", expected, actual)
	}
}
