package ransomnote

func canConstruct(ransomNote string, magazine string) bool {
	hashMap := map[rune]int{}

	for _, char := range magazine {
		hashMap[char]++
	}

	for _, char := range ransomNote {
		hashMap[char]--

		if hashMap[char] < 0 {
			return false
		}
	}
	return true
}
