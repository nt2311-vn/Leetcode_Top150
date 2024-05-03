package fixedpoint

func fixedPoint(arr []int) int {
	i, j := 0, len(arr)-1

	for i < j {
		mid := i + (j-i)/2

		if arr[mid] < mid {
			i = mid + 1
		} else {
			j = mid
		}
	}

	if arr[i] == i {
		return i
	}

	return -1
}
