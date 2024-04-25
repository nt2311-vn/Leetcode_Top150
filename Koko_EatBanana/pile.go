package kokoeatbanana

func minEatingSpeed(piles []int, h int) int {
	upper, lower, total := findBound(piles)

	for lower <= upper {
		mid := lower + (upper-lower)/2

		if total/mid <= h {
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}

	return lower
}

func findBound(piles []int) (int, int, int) {
	upper, lower, total := 0, 1, 0

	for _, pile := range piles {
		total += pile

		if pile > upper {
			upper = pile
		}

	}

	return upper, lower, total
}
