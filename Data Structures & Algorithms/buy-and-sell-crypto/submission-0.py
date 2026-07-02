class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        n = len(prices)
        buy_price = prices[0]
        mx = 0
        for i in range(n):
            if buy_price > prices[i]:
                buy_price = prices[i] 
            today_sell = prices[i] - buy_price
            if today_sell > mx:
                mx = today_sell
        return mx
        
        