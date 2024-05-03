package arrangingcoins

func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	i, j := 1, n

	for i < j {
		mid := i + (j-i)/2
		total := (mid * (mid + 1)) / 2

		if total <= n {
			i = mid + 1
		} else {
			j = mid
		}
	}

	return i - 1
}
