package validpalindrome

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	var cleaned strings.Builder

	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			cleaned.WriteRune(c)
		}
	}

	i, j := 0, cleaned.Len()-1

	for i < j {
		if cleaned.String()[i] != cleaned.String()[j] {
			return false
		}
		i++
		j--
	}

	return true
}
