package subsequence

//1312. 让字符串成为回文串的最少插入次数 https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/

//思路：超出最长的回文子串，剩下不是回文子串部分需要插入的字符，即 len(s) - len(LPS)
func minInsertions(s string) int {
	return len(s) - longestPalindromeSubseq(s)
}
