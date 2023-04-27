package dynamic_programming

import (
	"math"
)

//279. 完全平方数 https://leetcode.cn/problems/perfect-squares/
//递归 + dp table 的方式解决
func numSquares(n int) int {
	memo := make(map[int]int)

	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 {
			return 0
		}
		if _, ok := memo[n]; ok {
			return memo[n]
		}
		res := math.MaxInt32
		for i := 1; i*i <= n; i++ {
			sonSum := dp(n - i*i)
			if res > sonSum+1 {
				res = sonSum + 1
			}
		}
		memo[n] = res
		return memo[n]
	}
	return dp(n)
}

//迭代
func numSquaresRange(n int) int {
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = n + 1
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j*j <= i; j++ {
			res := dp[i-j*j] + 1
			if dp[i] > res {
				dp[i] = res
			}
		}
	}

	return dp[n]
}
