package subsequence

//712. 两个字符串的最小ASCII删除和  https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings/
//自底向上迭代的动态规划方式
func minimumDeleteSum(s1 string, s2 string) int {
	memo := make([][]int, len(s1)+1)
	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s2); i++ {
		memo[0][i] = memo[0][i-1] + int(s2[i-1])
	}

	for i := 1; i <= len(s1); i++ {
		memo[i][0] = memo[i-1][0] + int(s1[i-1])
	}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				memo[i+1][j+1] = memo[i][j]
			} else {
				if memo[i+1][j]+int(s2[j]) < memo[i][j+1]+int(s1[i]) {
					memo[i+1][j+1] = memo[i+1][j] + int(s2[j])
				} else {
					memo[i+1][j+1] = memo[i][j+1] + int(s1[i])
				}
			}
		}
	}

	return memo[len(s1)][len(s2)]
}

//自顶向下➕备忘录递归的动态规划方式的解法
func minimumDeleteSumDG(s1 string, s2 string) int {
	memo := make([][]int, len(s1)+1)
	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var traverse func(text1 string, i int, text2 string, j int) int
	traverse = func(text1 string, i int, text2 string, j int) int {
		//base case
		if i == len(text1) {
			res := 0
			for p := j; p < len(text2); p++ {
				res += int(text2[p])
			}
			memo[i][j] = res
		}

		if j == len(text2) {
			res := 0
			for p := i; p < len(text1); p++ {
				res += int(text1[p])
			}
			memo[i][j] = res
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if text1[i] == text2[j] {
			memo[i][j] = traverse(text1, i+1, text2, j+1)
		} else {
			l1 := traverse(text1, i+1, text2, j) + int(text1[i])
			l2 := traverse(text1, i, text2, j+1) + int(text2[j])
			if l1 < l2 {
				memo[i][j] = l1
			} else {
				memo[i][j] = l2
			}
		}
		return memo[i][j]
	}

	return traverse(s1, 0, s2, 0)
}
