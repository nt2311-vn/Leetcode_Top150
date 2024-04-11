package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)

	if len(pattern) != len(words) {
		return false
	}

	pattentTable := map[byte]string{}
	wordTable := map[string]byte{}

	for i := 0; i < len(pattern); i++ {
		wordVal, exist := pattentTable[pattern[i]]

		if exist && wordVal != words[i] {
			return false
		}
		pattentTable[pattern[i]] = words[i]

		patternVal, exist := wordTable[words[i]]

		if exist && patternVal != pattern[i] {
			return false
		}

		wordTable[words[i]] = pattern[i]
	}

	return true
}
