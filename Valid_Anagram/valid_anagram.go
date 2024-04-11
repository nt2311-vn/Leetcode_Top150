package validanagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	frequentMap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		frequentMap[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		frequentMap[t[i]]--

		if frequentMap[t[i]] < 0 {
			return false
		}
	}

	return true
}
