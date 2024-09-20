package main

func lengthOfLongestSubstring(s string) int {
	length := 0
	count := make([]int, 256)

	for i, j := 0, 0; i < len(s); i++ {
		count[s[i]]++
		for count[s[i]] > 1 {
			count[s[j]]--
			j++
		}
		if i-j+1 > length {
			length = i - j + 1
		}
	}

	return length
}
