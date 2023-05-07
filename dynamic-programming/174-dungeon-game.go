package dynamic_programming

import (
	"github.com/heyujiang/algorithm"
	"math"
)

//174. 地下城游戏 https://leetcode.cn/problems/dungeon-game/
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[m-1][n], dp[m][n-1] = 1, 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = algorithm.Min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			if dp[i][j] <= 0 {
				dp[i][j] = 1
			}
		}
	}

	return dp[0][0]
}

func calculateMinimumHPDG(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(dungeon [][]int, i, j int) int
	dp = func(dungeon [][]int, i, j int) int {
		if i == m-1 && j == n-1 {
			if dungeon[i][j] > 0 {
				return 1
			} else {
				return 1 - dungeon[i][j]
			}
		}

		if i == m {
			return math.MaxInt32
		}
		if j == n {
			return math.MaxInt32
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		res := algorithm.Min(dp(dungeon, i+1, j), dp(dungeon, i, j+1))
		if res-dungeon[i][j] <= 0 {
			memo[i][j] = 1
			return 1
		} else {
			memo[i][j] = res - dungeon[i][j]
			return memo[i][j]
		}
	}

	return dp(dungeon, 0, 0)
}
