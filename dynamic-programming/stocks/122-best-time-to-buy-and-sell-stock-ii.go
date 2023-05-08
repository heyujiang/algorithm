package stocks

import (
	"github.com/heyujiang/algorithm"
	"math"
)

//122. 买卖股票的最佳时机 II   https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
/**
Kw无限制：
dp[i][k][0] = max(dp[i-1][k][0],dp[i][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1],dp[i][k-1][0] - prices[i])

因为k不限制，所以k对状态转移方程没有影响，可以简化转移方程为：
dp[i][0] = max(dp[i-1][0],dp[i][1] + prices[i])
dp[i][1] = max(dp[i-1][1],dp[i][0] - prices[i])

当 i == 0 时：
dp[0][0] = max(dp[-1][0],dp[-1][1] + prices[i])
         = max(0,-∞+prices[i])
         = 0
dp[0][1] = max(dp[-1][1],dp[-1][0] - prices[i])
         = max(-∞,0-prices[i])
		 = -prices[i]
*/
func maxProfitII(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}

		dp[i][0] = algorithm.Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = algorithm.Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func maxProfitIIO1(prices []int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt32
	for i := 0; i < n; i++ {
		dp_i_0 = algorithm.Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = algorithm.Max(dp_i_1, dp_i_0-prices[i])
	}
	return dp_i_0
}
