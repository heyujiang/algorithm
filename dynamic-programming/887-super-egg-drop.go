package dynamic_programming

//887. 鸡蛋掉落  https://leetcode.cn/problems/super-egg-drop/
func superEggDrop(K int, N int) int {
	// m 最多不会超过 N 次（线性扫描）
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}
	// base case:
	// dp[0][..] = 0
	// dp[..][0] = 0
	// Go 默认初始化数组都为 0
	m := 0
	for dp[K][m] < N {
		m++
		for k := 1; k <= K; k++ {
			dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1
		}
	}
	return m
}
