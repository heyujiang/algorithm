package dynamic_programming

import "math"

//322.零钱兑换 https://leetcode.cn/problems/coin-change/submissions/
//剑指 Offer II 103. 最少的硬币数目 https://leetcode.cn/problems/gaM7Ch
//递归解法
func coinChange(coins []int, amount int) int {
	var dp func(amount int) int
	memo := make(map[int]int)
	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}

		if _, ok := memo[amount]; ok {
			return memo[amount]
		}

		res := math.MaxInt32
		for _, coin := range coins {
			sonNum := dp(amount - coin)
			if sonNum == -1 {
				continue
			}
			if res > sonNum+1 {
				res = sonNum + 1
			}
		}

		if res == math.MaxInt32 {
			memo[amount] = -1
		} else {
			memo[amount] = res
		}

		return memo[amount]
	}

	return dp(amount)
}

//迭代解法
func coinChangeRange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
