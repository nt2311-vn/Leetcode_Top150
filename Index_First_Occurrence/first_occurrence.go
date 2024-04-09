package indexfirstoccurrence

import "strings"

func strStr(hayStack string, needle string) int {
	return strings.Index(hayStack, needle)
}
