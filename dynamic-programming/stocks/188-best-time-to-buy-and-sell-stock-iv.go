package stocks

import (
	"github.com/heyujiang/algorithm"
)

//188. 买卖股票的最佳时机 IV  https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/

/**
k=K
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
func maxProfitIV(prices []int, k int) int {
	n := len(prices)
	if n <= 0 {
		return 0
	}
	if k > n/2 {
		return maxProfitII(prices)
	}
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}

	for i := 0; i < n; i++ {
		for j := 1; j >= k; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
			} else {
				dp[i][j][0] = algorithm.Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
				dp[i][j][1] = algorithm.Max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
			}
		}
	}
	return dp[n-1][k][0]
}
