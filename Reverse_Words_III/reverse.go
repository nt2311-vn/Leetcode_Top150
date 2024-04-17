package reversewordsiii

import "strings"

func reverse(s string) string {
	revWord := []byte(s)
	for i, j := 0, len(revWord)-1; i < j; i, j = i+1, j-1 {
		revWord[i], revWord[j] = revWord[j], revWord[i]
	}

	return string(revWord)
}

func reverseWords(s string) string {
	words := strings.Fields(s)

	for i, word := range words {
		revWord := reverse(word)
		words[i] = revWord
	}

	return strings.Join(words, " ")
}
