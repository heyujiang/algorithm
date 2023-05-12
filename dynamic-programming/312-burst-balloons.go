package dynamic_programming

import "github.com/heyujiang/algorithm"

//312. 戳气球 https://leetcode.cn/problems/burst-balloons/
func maxCoins(nums []int) int {
	n := len(nums)
	// 添加两侧的虚拟气球
	points := make([]int, n+2)
	points[0], points[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		points[i] = nums[i-1]
	}
	// base case 已经都被初始化为 0
	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+2)
	}
	// 开始状态转移
	// i 应该从下往上
	for i := n; i >= 0; i-- {
		// j 应该从左往右
		for j := i + 1; j < n+2; j++ {
			// 最后戳破的气球是哪个？
			for k := i + 1; k < j; k++ {
				// 择优做选择
				dp[i][j] = algorithm.Max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}
	return dp[0][n+1]
}
