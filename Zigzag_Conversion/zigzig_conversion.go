package zigzagconversion

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	arr := make([]string, numRows)
	check := true
	count := 0

	for _, ch := range s {
		arr[count] += string(ch)

		if check {
			count++
		} else {
			count--
		}

		if count == numRows-1 || count == 0 {
			check = !check
		}
	}

	return strings.Join(arr, "")
}
