package subsequence

import "github.com/heyujiang/algorithm"

//516. 最长回文子序列 https://leetcode.cn/problems/longest-palindromic-subsequence/
//自底向上迭代方式的动态规划
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = algorithm.Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

//自顶向下+备忘录的递归的动态规划
func longestPalindromeSubseqDG(s string) int {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(s))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dp func(s string, left, right int) int
	dp = func(s string, left, right int) int {
		//base case >>
		if left > right {
			return 0
		} else if left == right {
			return 1
		}
		//<< base case

		if memo[left][right] != -1 {
			return memo[left][right]
		}

		if s[left] == s[right] {
			memo[left][right] = dp(s, left+1, right-1) + 2
		} else {
			memo[left][right] = algorithm.Max(dp(s, left+1, right), dp(s, left, right-1))
		}
		return memo[left][right]
	}
	return dp(s, 0, len(s)-1)
}
