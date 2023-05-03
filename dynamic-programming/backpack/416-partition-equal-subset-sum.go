package backpack

//416. 分割等和子集 https://leetcode.cn/problems/partition-equal-subset-sum/
//剑指 Offer II 101. 分割等和子集 https://leetcode.cn/problems/NUPfPr/
/**
转化为背包问题：给一个可装载重量为sum/2的背包和N个物品，每个物品重量为nums[i]。现在要装物品，是否存在一种装法能够恰好将背包装满？
dp[i][j]bool 前i个物品，当前背包容量为j时，是否可以恰好将背包装满
**/
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	n := len(nums)
	sum = sum / 2

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
		dp[i][0] = true //背包容量没有了，即表示装满了
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 { //背包容量不足 不能装入
				dp[i][j] = dp[i-1][j]
			} else { //装入或者不装入
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[n][sum]
}
