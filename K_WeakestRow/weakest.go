package kweakestrow

import "sort"

type RowInfo struct {
	rowIndex     int
	countSoldier int
}

func kWeakestRows(mat [][]int, k int) []int {
	rowInfo := []RowInfo{}

	for index, row := range mat {
		rowInfo = append(rowInfo, RowInfo{rowIndex: index, countSoldier: binarySearch(row)})
	}

	sort.Slice(
		rowInfo,
		func(i, j int) bool {
			if rowInfo[i].countSoldier == rowInfo[j].countSoldier {
				return rowInfo[i].rowIndex < rowInfo[j].rowIndex
			}
			return rowInfo[i].countSoldier < rowInfo[j].countSoldier
		},
	)

	result := []int{}

	for i := 0; i < k; i++ {
		result = append(result, rowInfo[i].rowIndex)
	}

	return result
}

func binarySearch(row []int) int {
	i, j := 0, len(row)-1

	for i <= j {
		mid := int((i + j) / 2)

		if row[mid] == 1 {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return i
}
