package stocks

import "github.com/heyujiang/algorithm"

//309. 最佳买卖股票时机含冷冻期 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func maxProfitCooldown(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}

		if i-2 == -1 {
			dp[i][0] = algorithm.Max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = algorithm.Max(dp[i-1][1], -prices[i])
			continue
		}

		dp[i][0] = algorithm.Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = algorithm.Max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[n-1][0]
}
