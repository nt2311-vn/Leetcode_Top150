package isomorphicstrings

func isIsomorphic(s string, t string) bool {
	hashMap := map[rune]byte{}
	for index, char := range s {
		hashVal, exist := hashMap[char]

		if exist {
			if hashVal != t[index] {
				return false
			}
		} else {
			hashMap[char] = t[index]

			if len(hashMap) > 0 {
				for k, v := range hashMap {
					if v == t[index] && k != char {
						return false
					}
				}
			}
		}
	}

	return true
}
