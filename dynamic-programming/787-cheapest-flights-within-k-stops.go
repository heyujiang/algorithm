package dynamic_programming

import (
	"github.com/heyujiang/algorithm"
	"math"
)

//787. K 站中转内最便宜的航班 https://leetcode.cn/problems/cheapest-flights-within-k-stops/
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	indegree := make(map[int][][]int)

	for _, v := range flights {
		indegree[v[1]] = append(indegree[v[1]], []int{v[0], v[2]})
	}

	k = k + 1
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := range memo[i] {
			memo[i][j] = -2
		}
	}
	//从 src 出发，k步可以到达dst 的最小成本
	var dp func(dst int, k int) int
	dp = func(dst int, k int) int {
		if dst == src {
			return 0
		}
		if k == 0 {
			return -1
		}

		if memo[dst][k] != -2 {
			return memo[dst][k]
		}

		res := math.MaxInt32
		for _, v := range indegree[dst] {
			prev := dp(v[0], k-1)
			if prev == -1 {
				continue
			}
			res = algorithm.Min(res, prev+v[1])
		}

		if res == math.MaxInt32 {
			memo[dst][k] = -1
		} else {
			memo[dst][k] = res
		}
		return memo[dst][k]
	}

	return dp(dst, k)
}
