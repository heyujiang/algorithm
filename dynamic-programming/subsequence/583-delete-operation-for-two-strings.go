package subsequence

//583. 两个字符串的删除操作  https://leetcode.cn/problems/delete-operation-for-two-strings/

//思路：使word1和word2相同的最小步数使保留两个字符串的最长公共子串 ， 先求最长公共子串的长度
func minDistance(word1 string, word2 string) int {
	l := longestCommonSubsequence(word1, word2)
	return len(word1) + len(word2) - 2*l
}
