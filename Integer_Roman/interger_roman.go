package integerroman

var mapNum = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

var sliceInt = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	resultStr := ""

	for num > 0 {
		for _, numInt := range sliceInt {
			wholeVal := int(num / numInt)
			romanStr, _ := mapNum[numInt]
			num -= wholeVal * numInt
			for wholeVal > 0 {
				resultStr += romanStr
				wholeVal--
			}
		}
	}
	return resultStr
}
