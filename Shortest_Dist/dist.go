package shortestdist

func shortestToChar(s string, c byte) []int {
	result := []int{}

	for k := 0; k < len(s); k++ {
		if s[k] == c {
			result = append(result, 0)
		} else {
			i, j := k-1, k+1

			for i >= 0 && s[i] != c {
				i--
			}

			for j < len(s) && s[j] != c {
				j++
			}

			minDist := getMin(k-i, j-k)
			if i == -1 {
				minDist = j - k
			}

			if j == len(s) {
				minDist = k - i
			}
			result = append(result, minDist)
		}
	}

	return result
}

func getMin(a, b int) int {
	if a > b {
		return b
	}

	return a
}
