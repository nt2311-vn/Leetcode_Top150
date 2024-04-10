package containersmostwater

func maxNum(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func maxArea(height []int) int {
	left, right, currentMax := 0, len(height)-1, 0

	for left < right {
		if height[left] > height[right] {
			currentMax = maxNum(currentMax, (right-left)*height[right])
			right--
		} else {
			currentMax = maxNum(currentMax, (right-left)*height[left])
			left++
		}
	}

	return currentMax
}
