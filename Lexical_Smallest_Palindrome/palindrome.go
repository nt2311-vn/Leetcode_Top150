package lexicalsmallestpalindrome

func min(a, b byte) byte {
	if a < b {
		return a
	}

	return b
}

func makeSmallestPalindrome(s string) string {
	charArr := []byte(s)

	for i, j := 0, len(charArr)-1; i < j; i, j = i+1, j-1 {
		if charArr[i] != charArr[j] {
			charArr[i], charArr[j] = min(charArr[i], charArr[j]), min(charArr[i], charArr[j])
		}
	}

	return string(charArr)
}
