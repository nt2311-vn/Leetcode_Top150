package happynumber

import "strconv"

func isHappy(n int) bool {
	memoir := map[int]bool{}
	for n != 1 {

		if _, exist := memoir[n]; exist {
			return false
		}

		memoir[n] = true
		numStr := strconv.Itoa(n)
		result := 0

		for i := 0; i < len(numStr); i++ {
			charToInt, _ := strconv.Atoi(string(numStr[i]))
			result += charToInt * charToInt
		}

		if result == 1 {
			return true
		}

		n = result
	}

	return true
}
