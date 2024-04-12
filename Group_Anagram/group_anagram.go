package groupanagram

import (
	"sort"
)

func keyString(s string) string {
	result := ""
	runeSlice := []byte(s)
	sort.Slice(runeSlice, func(i, j int) bool { return runeSlice[i] < runeSlice[j] })

	for _, char := range runeSlice {
		result += string(char)
	}

	return result
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	anagramGroup := map[string][]string{}

	for _, str := range strs {
		keyStr := keyString(str)
		anagramGroup[keyStr] = append(anagramGroup[keyStr], str)
	}

	for _, value := range anagramGroup {
		result = append(result, value)
	}

	return result
}
