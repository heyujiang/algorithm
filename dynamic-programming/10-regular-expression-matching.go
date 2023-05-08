package dynamic_programming

import "fmt"

//10.正则表达式 https://leetcode.cn/problems/regular-expression-matching/
func isMatch(s string, p string) bool {
	memo := make(map[string]bool)
	sl := len(s)
	pl := len(p)
	var dp func(s string, i int, p string, j int) bool
	dp = func(s string, i int, p string, j int) bool {
		if j == pl {
			return i == sl
		}

		if i == sl { //检查 p[i:] 是否为 x*y*z* 形式
			if (pl-j)%2 == 1 {
				return false
			}
			for j = j + 1; j < pl; j = j + 2 {
				if p[j] != '*' {
					return false
				}
			}
			return true
		}

		key := fmt.Sprintf("%d,%d", i, j)
		if _, ok := memo[key]; ok {
			return memo[key]
		}

		if s[i] == p[j] || p[j] == '.' {
			if j < pl-1 && p[j+1] == '*' { //可以通配0次或者多次
				memo[key] = dp(s, i, p, j+2) || dp(s, i+1, p, j)
			} else { //匹配一次，匹配成功
				memo[key] = dp(s, i+1, p, j+1)
			}
		} else {
			if j < pl-1 && p[j+1] == '*' { //可以匹配0次
				memo[key] = dp(s, i, p, j+2)
			} else { //匹配失败
				memo[key] = false
			}
		}
		return memo[key]
	}

	return dp(s, 0, p, 0)
}
