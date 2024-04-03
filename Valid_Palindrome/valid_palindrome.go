package validpalindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	re := regexp.MustCompile(`/^[a-z]+$/i`)

	regexStr := re.FindAllString(strings.ToLower(s), -1)

	for i := 0; i <= len(regexStr)/2; i++ {
		if regexStr[i] != regexStr[len(regexStr)-1-i] {
			return false
		}
	}

	return true
}
