package subarraysum

import "testing"

func TestEdgeCase(t *testing.T) {
	nums := []int{4}
	k := 4
	expected := 1

	actual := minSubArrayLen(k, nums)

	if actual != expected {
		t.Errorf("Failed edge case, expected: %d but got %d\n", expected, actual)
	}
}
