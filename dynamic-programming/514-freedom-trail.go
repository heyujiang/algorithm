package dynamic_programming

import (
	"fmt"
	"github.com/heyujiang/algorithm"
)

//514. 自由之路  https://leetcode.cn/problems/freedom-trail/
func findRotateSteps(ring string, key string) int {
	n := len(ring)
	m := len(key)

	memo := make(map[string]int)

	var dp func(ring string, i int, key string, j int, isZ bool) int
	dp = func(ring string, i int, key string, j int, isZ bool) int {
		if j == m {
			return 0
		}

		k := fmt.Sprintf("%d:%v:%d", i, isZ, j)
		if _, ok := memo[k]; ok {
			return memo[k]
		}

		if ring[i] == key[j] { //相同的时候才正 or 倒
			memo[k] = algorithm.Min(dp(ring, i, key, j+1, isZ), dp(ring, i, key, j+1, !isZ)) + 1
		} else { //不同的时候只能一直正或倒
			if isZ {
				i = i + 1
				if i == n {
					i = 0
				}
				memo[k] = dp(ring, i, key, j, isZ) + 1
			} else {
				i = i - 1
				if i == -1 {
					i = n - 1
				}
				memo[k] = dp(ring, i, key, j, isZ) + 1
			}
		}
		return memo[k]
	}

	return algorithm.Min(dp(ring, 0, key, 0, true), dp(ring, 0, key, 0, false))
}
