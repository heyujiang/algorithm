package greed

import "github.com/heyujiang/algorithm"

//45. 跳跃游戏II  https://leetcode.cn/problems/jump-game-ii/
func jump(nums []int) int {
	n := len(nums)
	end, farthest, jumps := 0, 0, 0
	for i := 0; i < n-1; i++ {
		farthest = algorithm.Max(nums[i]+i, farthest)
		if end == i {
			jumps++
			end = farthest
		}
	}
	return jumps
}
