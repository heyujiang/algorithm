package subsequence

import "sort"

//254.俄罗斯套娃信封问题  https://leetcode.cn/problems/russian-doll-envelopes/
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	//dp := make([]int,len(envelopes))  //动态规划解法 O(n^2)
	//res := 0
	//for i:=0;i<len(envelopes);i++{
	//	dp[i] = 1
	//	for j:=0;j<i;j++{
	//		if envelopes[i][1] > envelopes[j][1] {
	//			if dp[i] < dp[j] + 1 {
	//				dp[i] = dp[j]+1
	//			}
	//		}
	//	}
	//	if res < dp[i] {
	//		res = dp[i]
	//	}
	//}
	//return res

	//二分搜索解法  O(nlongn)
	top := make([]int, len(envelopes))
	piles := 0
	for i := 0; i < len(envelopes); i++ {
		left, right := 0, piles

		for left < right {
			mid := left + (right-left)/2

			if top[mid] < envelopes[i][1] {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == piles {
			piles++
		}

		top[left] = envelopes[i][1]
	}
	return piles
}
