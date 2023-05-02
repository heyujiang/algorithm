package dynamic_programming

//509. 斐波那契数(https://leetcode.cn/problems/fibonacci-number/)

//暴力穷举迭代方式
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

//带备忘录穷举迭代方式  自顶向下的递归解法
func fibMemo(n int) int {
	memo := make([]int, n+1)
	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		if memo[n] == 0 {
			memo[n] = dp(n-2) + dp(n-1)
		}
		return memo[n]
	}
	return dp(n)
}

//自底向上的迭代方式实现  空间复杂度O(N)
func fibFromBottom(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[1] = 1 //base case

	for i := 2; i < n+1; i++ { //状态转移
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

//还可以再进一步优化将空间复杂度降到O(1)
//如果发现每次转移只需要dp table中的一部分，那么就可以尝试缩小dp table的大小，只记录必要的数据，从而降低空间复杂度
func fibFromBottomO1(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	dp_0, dp_1 := 0, 1 //base case

	for i := 2; i < n+1; i++ { //状态转移
		dp := dp_0 + dp_1
		dp_0 = dp_1
		dp_1 = dp
	}
	return dp_1
}
