package stocks

import (
	"github.com/heyujiang/algorithm"
	"math"
)

//714. 买卖股票的最佳时机含手续费 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
func maxProfitFee(prices []int, fee int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt32
	for i := 0; i < n; i++ {
		dp_i_0 = algorithm.Max(dp_i_0, dp_i_1+prices[i]-fee)
		dp_i_1 = algorithm.Max(dp_i_1, dp_i_0-prices[i])
	}
	return dp_i_0
}
