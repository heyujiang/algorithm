package house_robber

import (
	"github.com/heyujiang/algorithm"
	"math"
)

//198. 打家劫舍  https://leetcode.cn/problems/house-robber/

/**
第一步：确定状态和选择	：
	状态就是有几个房屋：房屋数目是给定的数组长度；
	选择就是可偷可不偷：可用0表示不偷，1表示偷；

	dp[2][0]表示第二个房间，不偷窃；dp[3][1]表示第三个房间偷；

第二步确认状态转移方程：
	因为两个相邻的房间被偷会触发警报，所以如果第i个房间选择偷窃，那么第i-1个房间肯定不能偷窃；如果第i个房间不偷窃，那么第i-1个房间可偷可不偷；
	dp[i][0] = max(dp[i-1][0],dp[i][i-1][1])
	dp[i][1] = dp[i-1][0] + nums[i]

第三步确认base case：
	i - 1 = -1时，因为房间数是从0开始的：
		dp[-1][0] -1房间肯定不存在，偷盗的财物为0；
		dp[-1][1] -1房间不存在，那么肯定不能偷窃，又因为求的是偷盗财物的最大值，所以给设置为-∞；

	根据状态转移方程得到的实际base case：
	dp[0][0] = max(dp[-1][1],dp[-1][0])
         = max(-∞，0)
		 = 0

	dp[0][1] = dp[-1][0] + nums[i]
			 = 0 + nums[i]
			 = nums[i]

第四步确认结果取值：
	小偷肯定要光顾所有的房子，所以结果的状态肯定是最后一个房子，最后一个房子可以选择偷或者不偷,而为了偷窃金额最高，那么肯定选择最后一个房子偷或者不偷相对大的那个选择 ：
	所以结果是 max(dp[n][0],dp[n][1])
*/
func rob(nums []int) int {
	n := len(nums)
	dp := make([][2]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = nums[i]
			continue
		}

		dp[i][0] = algorithm.Max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return algorithm.Max(dp[n-1][0], dp[n-1][1])
}

//空间复杂度为O(1)的解法
func robO1(nums []int) int {
	dp_i_0, dp_i_1 := 0, math.MinInt32
	for i := 0; i < len(nums); i++ {
		dp_i_0, dp_i_1 = algorithm.Max(dp_i_0, dp_i_1), dp_i_0+nums[i]
	}

	return algorithm.Max(dp_i_0, dp_i_1)
}
