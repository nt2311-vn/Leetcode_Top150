package minarrowburstballoon

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinArrowShots(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}

	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })

	count := 1
	end := points[0][1]

	for _, point := range points {
		if point[0] <= end {
			end = min(end, point[1])
		} else {
			count++
			end = point[1]
		}
	}

	return count
}
