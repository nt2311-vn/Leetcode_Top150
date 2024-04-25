package intersection3array

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	result := []int{}
	hashMap := map[int]bool{}

	for i, j, k := 0, 0, 0; i < len(arr1) && j < len(arr2) && k < len(arr3); {
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			if _, exist := hashMap[arr1[i]]; !exist {
				result = append(result, arr1[i])
				hashMap[arr1[i]] = true
			}
			i, j, k = i+1, j+1, k+1
		} else if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr3[k] {
			j++
		} else {
			k++
		}
	}

	return result
}
