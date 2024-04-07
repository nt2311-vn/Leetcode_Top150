package insertionsort

type Pair struct {
	key   int
	value string
}

func insertionSort(pairs []Pair) [][]Pair {
	results := [][]Pair{}

	for i := 0; i < len(results); i++ {
		j := i - 1

		for j >= 0 && pairs[j+1].key < pairs[j].key {
			pairs[j+1], pairs[j] = pairs[j], pairs[j+1]
		}
		results = append(results, pairs)
	}

	return results
}
