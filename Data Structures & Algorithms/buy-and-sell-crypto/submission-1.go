func maxProfit(prices []int) int {
	cur := prices[0]
	mx := 0
	for _, v := range prices {
		if cur > v {
			cur = v
		} else {
			mx = max(mx, v - cur)
		}
	}
	return mx
}
