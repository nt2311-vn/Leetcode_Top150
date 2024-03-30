package baseballgame

import (
	"strconv"
)

func callPoints(operations []string) int {
	scores := []int{}

	for _, v := range operations {
		val, err := strconv.Atoi(v)

		if err == nil {
			scores = append(scores, val)
		} else {
			switch v {
			case "+":
				newRecord := scores[len(scores)-1] + scores[len(scores)-2]
				scores = append(scores, newRecord)
				break

			case "D":
				newRecord := scores[len(scores)-1] * 2
				scores = append(scores, newRecord)
				break

			case "C":
				scores = scores[:len(scores)-1]
				break

			default:
				panic("Operation not defined")
			}
		}
	}

	result := 0

	for _, v := range scores {
		result += v
	}

	return result
}
