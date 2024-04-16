package crawlerlogfolder

// var mapCd = map[string]int{"../": -1, "./": 0}

func minOperations(logs []string) int {
	// minReturn := 0
	//
	// for _, operation := range logs {
	// 	val, exist := mapCd[operation]
	//
	// 	if !exist {
	// 		minReturn++
	// 	} else {
	// 		minReturn += val
	// 		if minReturn < 0 {
	// 			minReturn = 0
	// 		}
	// 	}
	// }
	//
	// return minReturn

	stack := []byte{}

	for _, operation := range logs {
		if operation == "../" {
			if len(stack) > 0 {
				stack = stack[1:]
			}
		} else if operation == "./" {
			continue
		} else {
			stack = append(stack, 1)
		}
	}

	return len(stack)
}
