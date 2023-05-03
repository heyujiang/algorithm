package backpack

//518. 零钱兑换 II   https://leetcode.cn/problems/coin-change-ii/
/**
第一步：确认状态和选择 ：
状态有两个，一个是背包的容量和可选择的物品，选择就是装入背包或者不装入背包；这里背包的容量就是总金额，可选择的物品就是硬币数组coins；

第二步：明确dp数组的定义：
dp[i][j] = x 是指在可选择物品的前i个在用量为j时，有x种方式装满背包
这题是在可选硬币数组的前i个硬币种，若想凑齐金额j，有x种凑法；

第三部：根据选择，思考状态转移的逻辑
如果不把第i个物品装入背包，也就是使用coins[i-1]这枚硬币，那么凑出j金额的方式dp[i][j] 等于 dp[i-1][j].
如果把第i个物品放入背包，也就是使用coins[i-1]这枚硬币，有因为硬币不限，那么凑出j金额的方式dp[i][j]应该等于dp[i][j-coins[i-1]].
*/
func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i-1] { //硬币的额度大于需要的金额 只能不放入
				dp[i][j] = dp[i-1][j]
			} else { //可以放入 可以不放入
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}
