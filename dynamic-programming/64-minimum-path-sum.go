package dynamic_programming

import "github.com/heyujiang/algorithm"

//64. 最小路径和  https://leetcode.cn/problems/minimum-path-sum/
//剑指 Offer II 099. 最小路径之和 https://leetcode.cn/problems/0i0mDW/submissions/
func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = algorithm.Min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[n-1][m-1]
}
