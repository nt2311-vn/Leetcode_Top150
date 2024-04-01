package besttimebuysell

import "testing"

func TestFirstCase(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5

	actual := maxProfit(prices)
	if expected != actual {
		t.Errorf("Failed first test: expected %d but got %d", expected, actual)
	}
}

func TestSecondCase(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expected := 0

	actual := maxProfit(prices)
	if expected != actual {
		t.Errorf("Failed second test: expected %d but got %d", expected, actual)
	}
}
