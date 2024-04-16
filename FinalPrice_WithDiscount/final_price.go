package finalpricewithdiscount

func finalPrices(prices []int) []int {
	result := make([]int, len(prices))

	for i := 0; i < len(prices); i++ {
		discount := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				discount = prices[j]
				break
			}
		}

		result[i] = prices[i] - discount
	}

	return result
}
