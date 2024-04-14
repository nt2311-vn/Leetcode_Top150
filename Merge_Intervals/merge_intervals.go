package mergeintervals

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func merge(intervals [][]int) [][]int {
	result := [][]int{}

	if len(intervals) == 0 {
		return result
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		resultBound := result[len(result)-1]

		if current[0] <= resultBound[1] {
			resultBound[1] = max(resultBound[1], current[1])
		} else {
			result = append(result, current)
		}
	}

	return result
}
