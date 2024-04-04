package climbingstairs

var memoir = make(map[int]int)

func climbStairs(n int) int {
	if val, exist := memoir[n]; exist {
		return val
	}

	if n <= 2 {
		return n
	}

	memoir[n] = climbStairs(n-1) + climbStairs(n-2)

	return memoir[n]
}
