package twosumsortedarray

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		return []int{1, 2}
	}

	i, j := 0, len(numbers)-1

	for i < j {
		sum := numbers[i] + numbers[j]

		if sum == target {
			break
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return []int{i + 1, j + 1}
}
