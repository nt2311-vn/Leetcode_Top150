package besttimebuysellii

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	buyPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else {
			profit := prices[i] - buyPrice

			if profit > 0 {
				maxProfit += profit
				buyPrice = prices[i]
			}
		}
	}

	return maxProfit
}
