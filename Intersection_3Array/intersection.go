package intersection3array

import "sort"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	/*
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
	*/

	result := []int{}
	firstSearch := []int{}

	for _, num := range arr1 {
		if binarySearch(num, arr2) {
			firstSearch = append(firstSearch, num)
		}
	}

	for _, num := range arr3 {
		if binarySearch(num, firstSearch) {
			result = append(result, num)
		}
	}

	sort.Ints(result)

	return result
}

func binarySearch(val int, target []int) bool {
	start, end := 0, len(target)-1

	for start <= end {
		mid := (start + end) / 2

		if target[mid] < val {
			start = mid + 1
		} else if target[mid] > val {
			end = mid - 1
		} else {
			return true
		}
	}

	return false
}
