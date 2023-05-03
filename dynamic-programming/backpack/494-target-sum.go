package backpack

import "strconv"

//494. 目标和  https://leetcode.cn/problems/target-sum/
//剑指 Offer II 102. 加减的目标值  https://leetcode.cn/problems/YaVDxD/

//自顶向下递归的方式解决
func findTargetSumWays(nums []int, target int) int {
	var dp func(nums []int, i int, target int) int
	dp = func(nums []int, i int, target int) int {
		if i >= len(nums) {
			if target == 0 {
				return 1
			} else {
				return 0
			}
		}

		return dp(nums, i+1, target-nums[i]) + dp(nums, i+1, target+nums[i])
	}
	return dp(nums, 0, target)
}

//自顶向下+备忘录递归的动态规划的方式解决
func findTargetSumWaysMemo(nums []int, target int) int {
	memo := make(map[string]int)
	var dp func(nums []int, i int, target int) int
	dp = func(nums []int, i int, target int) int {
		if i >= len(nums) {
			if target == 0 {
				return 1
			} else {
				return 0
			}
		}

		key := strconv.Itoa(i) + ":" + strconv.Itoa(target)
		if _, ok := memo[key]; ok {
			return memo[key]
		}

		memo[key] = dp(nums, i+1, target-nums[i]) + dp(nums, i+1, target+nums[i])
		return memo[key]
	}
	return dp(nums, 0, target)
}

//可以转换为背包问题
/**
添加两种符号那就是正负，那么就是:
Sum(A) - Sum(B) = target
Sum(A) - target = Sum(B)
Sum(A) + Sum(B) - target =  Sum(B) + Sum(B)
Sum(nums) - target = 2 * Sum(B)
(Sum(nums) - target)/2 =  Sum(B)
即转为了背包问题：
又一个背包，容量为sum，现在有N个物品，第i个物品的重量为nums[i-1](注意1<=i<=N)，每个物品只有一个，请问有几种不同的方法能够恰好装满这个背包

第一步需要明确状态和选择，状态就是背包的容量和可选择的物品，状态就是可装可不装

第二步明确dp数组的定义：
dp[i][j]表示，若只在前i个物品中选择，若当前背包的容量为j，则最多有x种方法可以装满恰好背包

第三部，根据选择，思考状态转移的逻辑：
如果不把nums[i]不算入子集，或者说不放入背包，那么恰好装满背包的方法数取决于上一个状态dp[i-1][j]，继承之前的结果；
如果把nums[i]算入子集，或者说放入背包，那么只要看前i-1个物品有几种方法可以装满j-nums[i-1]的重量就行，所以取决于dp[i-1][j-nums[i-1]]

由于dp[i][j]为装满背包的总方法数，所以应该是上述两种结果的求和，得到状态转移方程为：
dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
*/
func findTargetSumWaysRange(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum < target || (sum-target)%2 != 0 {
		return 0
	}

	sum = (sum - target) / 2
	n := len(nums)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= sum; j++ {
			if j >= nums[i-1] { //空间充足，可装可不装
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else { //空间不足，只能不装
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][sum]
}
