package h_index

import "sort"

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	for i, v := range citations {
		position := len(citations) - i

		if position <= v {
			return position
		}
	}
	return 0
}
