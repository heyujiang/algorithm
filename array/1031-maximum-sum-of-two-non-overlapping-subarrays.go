package array

//1031. 两个非重叠子数组的最大和(https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays)
func maxSumTwoNoOverlap(nums []int, firstLen, secondLen int) (ans int) {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x // 计算 nums 的前缀和
	}
	maxSumA, maxSumB := 0, 0
	for i := firstLen + secondLen; i <= n; i++ {
		maxSumA = max(maxSumA, s[i-secondLen]-s[i-secondLen-firstLen])
		maxSumB = max(maxSumB, s[i-firstLen]-s[i-firstLen-secondLen])
		ans = max(ans, max(maxSumA+s[i]-s[i-secondLen], // 左 a 右 b
			maxSumB+s[i]-s[i-firstLen])) // 左 b 右 a
	}
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
