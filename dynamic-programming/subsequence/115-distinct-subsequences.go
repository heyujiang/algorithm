package subsequence

//115. 不同的子序列 https://leetcode.cn/problems/distinct-subsequences/
func numDistinct(s string, t string) int {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(t))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(s string, i int, t string, j int) int
	dp = func(s string, i int, t string, j int) int {
		if j >= len(t) {
			return 1
		}
		if len(s)-i < len(t)-j {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		res := 0
		for k := i; k < len(s); k++ { //从s串的不同位置开始寻找t串
			if s[k] == t[j] {
				res += dp(s, k+1, t, j+1)
			}
		}
		memo[i][j] = res
		return res
	}

	return dp(s, 0, t, 0)
}

func numDistinct2(s string, t string) int {
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(t))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(s string, i int, t string, j int) int
	dp = func(s string, i int, t string, j int) int {
		if j >= len(t) {
			return 1
		}
		if len(s)-i < len(t)-j {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		res := 0
		if s[i] == t[j] { //如果s[i] == t[j] , s[i]可以参与子序列也可以不参与子序列
			res += dp(s, i+1, t, j+1) + dp(s, i+1, t, j)
		} else { //s[i] != t[j] ，s[i]显然不属于子序列
			res += dp(s, i+1, t, j)
		}

		memo[i][j] = res
		return res
	}

	return dp(s, 0, t, 0)
}
