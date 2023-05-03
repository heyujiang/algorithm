package subsequence

import "github.com/heyujiang/algorithm"

//72. 编辑距离 https://leetcode.cn/problems/edit-distance/
// 自顶向下带备忘录的递归的动态规划
func minDistance2(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i := range memo {
		memo[i] = make([]int, len(word2))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dp func(word1 string, i int, word2 string, j int) int

	dp = func(word1 string, i int, word2 string, j int) int {
		if i >= len(word1) {
			return len(word2) - j
		}
		if j >= len(word2) {
			return len(word1) - i
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if word1[i] == word2[j] {
			memo[i][j] = dp(word1, i+1, word2, j+1)
		} else {
			memo[i][j] = algorithm.Min(dp(word1, i+1, word2, j+1)+1, algorithm.Min(dp(word1, i+1, word2, j)+1, dp(word1, i, word2, j+1)+1))
		}
		return memo[i][j]
	}

	return dp(word1, 0, word2, 0)
}

//自底向上的迭代的动态规划方式
func minDistance2Range(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := range dp {
		dp[i] = make([]int, len(word1)+1)
	}

	for i := 0; i <= len(word1); i++ {
		dp[0][i] = i
	}

	for i := 0; i <= len(word2); i++ {
		dp[i][0] = i
	}

	for i := 1; i <= len(word2); i++ {
		for j := 1; j <= len(word1); j++ {
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = algorithm.Min(dp[i-1][j]+1, algorithm.Min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}
	return dp[len(word2)][len(word1)]
}
