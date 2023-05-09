package house_robber

import (
	"github.com/heyujiang/algorithm"
	"math"
)

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
