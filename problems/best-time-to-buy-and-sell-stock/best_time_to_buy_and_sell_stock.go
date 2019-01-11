package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	ans, min := 0, 0
	for i, v := range prices {
		if v < prices[min] {
			min = i
		} else if v-prices[min] > ans {
			ans = v - prices[min]
		}
	}
	return ans
}
