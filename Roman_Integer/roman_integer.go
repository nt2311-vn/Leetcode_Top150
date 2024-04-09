package romaninteger

var (
	mapSingle = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	mapDouble = map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
)

func romanToInt(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			val, exist := mapDouble[s[i:i+2]]

			if exist {
				result += val
				i++
				continue
			}
		}
		result += mapSingle[s[i]]
	}

	return result
}
