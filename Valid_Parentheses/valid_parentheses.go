package validparentheses

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	mapComplement := map[rune]rune{'{': '}', '[': ']', '(': ')'}

	stack := []rune{}

	for _, v := range s {
		if len(stack) > 0 && v == mapComplement[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}

	return len(stack) == 0
}
