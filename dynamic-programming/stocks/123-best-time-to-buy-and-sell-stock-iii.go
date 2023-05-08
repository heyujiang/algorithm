package stocks

import (
	"github.com/heyujiang/algorithm"
)

//123. 买卖股票的最佳时机 III  https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
/**
k=2
dp[i][k][0] = max(dp[i-1][k][0],dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1],dp[i-1][k-1][0] - prices[i])

当 i == 0 时 ：
dp[0][k][0] = max(dp[-1][k][0],dp[-1][k][1] + prices[i])
            = max(0,-∞ + prices[i])
            = 0
dp[0][k][1] = max(dp[-1][k][1],dp[i-1][k-1][0] - prices[i])
			= max(-∞,0-prices[i])
			= -prices[i]

当 k = 0 时 ：
dp[i][0][0] = 0
dp[i][0][1] = -∞
*/
func maxProfitIII(prices []int) int {
	n := len(prices)
	dp := make([][][2]int, n+1)
	for i := range dp {
		dp[i] = make([][2]int, 3)
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= 2; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[0]
				continue
			}

			dp[i][j][0] = algorithm.Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = algorithm.Max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][2][0]
}
