package maxdepthnest

func maxDepth(s string) int {
	maxDepth, currentDepth := 0, 0

	for _, char := range s {
		if char == '(' {
			currentDepth++
			if currentDepth > maxDepth {
				maxDepth = currentDepth
			}
		}

		if char == ')' {
			currentDepth--
		}
	}

	return maxDepth
}
