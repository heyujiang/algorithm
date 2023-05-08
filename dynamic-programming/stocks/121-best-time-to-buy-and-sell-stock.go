package stocks

import (
	"github.com/heyujiang/algorithm"
	"math"
)

//121. 买卖股票的最佳时机   https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

//暴力解法 时间复杂度 O(N^2)
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit = algorithm.Max(profit, prices[j]-prices[i])
		}
	}
	return profit
}

//动态规划解法
func maxProfit2(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}

		dp[i][0] = algorithm.Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = algorithm.Max(dp[i-1][1], 0-prices[i])
	}
	return dp[n-1][0]
}

//空间优化版本，空间复杂度为O（1）
func maxProfit2O1(prices []int) int {
	n := len(prices)
	dp_i_0, dp_i_1 := 0, math.MinInt32
	for i := 0; i < n; i++ {
		dp_i_0 = algorithm.Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = algorithm.Max(dp_i_1, -prices[i])
	}
	return dp_i_0
}

func maxProfit3(prices []int) int {
	n := len(prices)
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-888, -888}
	}

	var dp func(i, status int) int
	dp = func(i, status int) int {
		if status == 0 { //卖出
			if i == 0 { //第一天没有可卖的
				return 0
			}
			if memo[i][0] != -888 {
				return memo[i][0]
			}
			memo[i][0] = algorithm.Max(dp(i-1, 0), dp(i-1, 1)+prices[i])
			return memo[i][0]
		} else { //买入
			if i == 0 { //第一天需要花费
				return -prices[i]
			}
			if memo[i][1] != -888 {
				return memo[i][1]
			}
			memo[i][1] = algorithm.Max(dp(i-1, 1), -prices[i])
			return memo[i][1]
		}
	}
	return dp(n-1, 0)
}
