package gas_station

func canCompleteCirCuit(gas, cost []int) int {
	fill, energy, start := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		fill += gas[i] - cost[i]
		energy += gas[i] - cost[i]

		if fill < 0 {
			start = i + 1
			fill = 0
		}
	}

	if energy < 0 {
		return -1
	}

	return start
}
