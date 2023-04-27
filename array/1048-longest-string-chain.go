package array

import "sort"

//1048. 最长字符串链 (https://leetcode.cn/problems/longest-string-chain/)
//
//现将字符串按长度从小到大排序
//将单词去掉一个字母即为此单词在词链中的前身 所以此单词的词链长度为前身的词链长度； 多个取最大的  （动态规划思想）
func longestStrChain(words []string) int {
	res := 0
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[i])
	})

	dp := make(map[string]int, len(words))
	for _, word := range words {
		dp[word] = 1
		for i := 0; i < len(word); i++ {
			nowWord := word[0:i] + word[i+1:]
			if num, ok := dp[nowWord]; ok {
				if dp[word] > num+1 {
					dp[word] = num + 1
				}
			}
		}
	}

	for _, v := range dp {
		if res < v {
			res = v
		}
	}
	return res
}
