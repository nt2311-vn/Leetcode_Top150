package removeoutermostparenthese

func removeOuterParentheses(s string) string {
	result := ""
	pointer := 1
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			pointer++
		} else {
			pointer--
		}

		if pointer == 0 {
			pointer = 1
			i++
		} else {
			result += string(s[i])
		}
	}

	return result
}
