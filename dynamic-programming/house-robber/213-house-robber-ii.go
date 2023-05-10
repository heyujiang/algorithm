package house_robber

import (
	"github.com/heyujiang/algorithm"
	"math"
)

//213. 打家劫舍II  https://leetcode.cn/problems/house-robber-ii/
//剑指 Offer II 090. 环形房屋偷盗  https://leetcode.cn/problems/PzWKhm/

//环形房屋，那么首尾相连，那么要么偷首不偷尾、要么偷尾不偷首、要么首尾都不偷 ； 所以比较三种方式的最大值即可，首尾都不偷肯定小于另位两种方式，所以可以不同考虑了
func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return algorithm.Max(robF(nums, 0, len(nums)-1), robF(nums, 1, len(nums)))
}

func robF(nums []int, start, end int) int {
	dp := make([][2]int, end)

	for i := start; i < end; i++ {
		if i == start {
			dp[i][0] = 0
			dp[i][1] = nums[i]
			continue
		}

		dp[i][0] = algorithm.Max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return algorithm.Max(dp[end-1][0], dp[end-1][1])
}

func robFO1(nums []int, start, end int) int {
	dp_i_0, dp_i_1 := 0, math.MinInt32
	for i := start; i < end; i++ {
		dp_i_0, dp_i_1 = algorithm.Max(dp_i_0, dp_i_1), dp_i_0+nums[i]
	}
	return algorithm.Max(dp_i_0, dp_i_1)
}
