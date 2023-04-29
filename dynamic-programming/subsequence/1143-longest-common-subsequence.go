package subsequence

//1143. 最长公共子序列 https://leetcode.cn/problems/longest-common-subsequence/
//剑指 Offer II 095. 最长公共子序列  https://leetcode.cn/problems/qJnOS7/submissions/
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
