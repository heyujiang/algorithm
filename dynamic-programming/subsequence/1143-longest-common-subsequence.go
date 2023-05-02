package subsequence

//1143. 最长公共子序列 https://leetcode.cn/problems/longest-common-subsequence/
//剑指 Offer II 095. 最长公共子序列  https://leetcode.cn/problems/qJnOS7/submissions/

//自底向上迭代的动态规划方式
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				if dp[i+1][j] > dp[i][j+1] {
					dp[i+1][j+1] = dp[i+1][j]
				} else {
					dp[i+1][j+1] = dp[i][j+1]
				}
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

// 自顶向下的带备忘录的递归的动态规划方式
func longestCommonSubsequenceDG(text1 string, text2 string) int {
	memo := make([][]int, len(text1)+1)
	for i := range memo {
		memo[i] = make([]int, len(text2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var traverse func(text1 string, i int, text2 string, j int) int
	traverse = func(text1 string, i int, text2 string, j int) int {
		//base case
		if i == len(text1) || j == len(text2) {
			memo[i][j] = 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if text1[i] == text2[j] {
			memo[i][j] = traverse(text1, i+1, text2, j+1) + 1
		} else {
			l1 := traverse(text1, i+1, text2, j)
			l2 := traverse(text1, i, text2, j+1)
			if l1 > l2 {
				memo[i][j] = l1
			} else {
				memo[i][j] = l2
			}
		}
		return memo[i][j]
	}

	return traverse(text1, 0, text2, 0)
}
