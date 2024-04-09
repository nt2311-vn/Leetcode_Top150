package lengthlastword

import "strings"

func lengthOfLastWord(s string) int {
	wordRange := strings.Fields(s)

	return len(wordRange[len(wordRange)-1])
}
