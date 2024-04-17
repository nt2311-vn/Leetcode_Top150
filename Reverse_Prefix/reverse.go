package reverseprefix

func reversePrefix(word string, ch byte) string {
	byteArr := []byte(word)
	isFound, index := false, 0
	for ; index < len(byteArr); index++ {
		if byteArr[index] == ch {
			isFound = true
			break
		}
	}

	if !isFound {
		return word
	}

	for i, j := 0, index; i < j; i, j = i+1, j-1 {
		byteArr[i], byteArr[j] = byteArr[j], byteArr[i]
	}

	return string(byteArr)
}
