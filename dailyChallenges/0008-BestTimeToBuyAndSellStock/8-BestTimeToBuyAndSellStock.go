func maxProfit(prices []int) int {
	l, r := 0, 1
	maxProfit, profit := 0, 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit = prices[r] - prices[l]
			maxProfit = max(maxProfit, profit)
		} else {
			l = r
		}
		r += 1
	}
	return maxProfit

}