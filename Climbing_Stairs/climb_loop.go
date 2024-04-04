package climbingstairs

func climbStairsLoop(n int) int {
	if n <= 2 {
		return n
	}

	first, second := 1, 2

	for i := 3; i <= n; i++ {
		third := first + second

		first, second = second, third
	}

	return second
}
