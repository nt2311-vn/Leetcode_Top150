package countnegative

unc countNegatives(grid [][]int) int {
	result := 0

	for _, arr := range grid {
		result += binarySearch(arr)
	}

	return result
}

func binarySearch(target []int) int {
	i, j := 0, len(target)-1

	for i <= j {
		mid := int((i + j) / 2)

		if target[mid] >= 0 {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return len(target[i:])
}
