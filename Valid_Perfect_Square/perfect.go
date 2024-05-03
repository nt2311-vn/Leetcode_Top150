package validperfectsquare

func isPerfectSquare(num int) bool {
	if num < 4 {
		if num == 1 {
			return true
		} else {
			return false
		}
	}

	i, j := 1, int(num/2)

	for i <= j {
		mid := i + (j-i)/2

		if mid*mid == num {
			return true
		} else if mid < num/mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return false
}
