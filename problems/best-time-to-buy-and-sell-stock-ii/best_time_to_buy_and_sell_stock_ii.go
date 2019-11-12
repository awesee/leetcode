package problem122

func maxProfit(prices []int) int {
	maxProfit := 0
	for i, p := range prices {
		if i > 0 && p > prices[i-1] {
			maxProfit += p - prices[i-1]
		}
	}
	return maxProfit
}
