package dynamic_programming

import (
	"github.com/heyujiang/algorithm"
	"math"
)

//931. 下降路径最小和   https://leetcode.cn/problems/minimum-falling-path-sum/
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+2)
		dp[i][0] = math.MaxInt32
		dp[i][m+1] = math.MaxInt32
	}

	for i := 1; i < m+2; i++ {
		dp[n][i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = algorithm.Min(algorithm.Min(dp[i-1][j-1], dp[i-1][j]), dp[i-1][j+1]) + matrix[i-1][j-1]
		}
	}

	res := math.MaxInt32
	for i := 1; i < m+2; i++ {
		if res > dp[n][i] {
			res = dp[n][i]
		}
	}
	return res
}
