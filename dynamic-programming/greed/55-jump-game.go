package greed

import "github.com/heyujiang/algorithm"

//55. 跳跃游戏  https://leetcode.cn/problems/jump-game/
func canJump(nums []int) bool {
	n := len(nums)
	farthest := 0
	for i := 0; i < n-1; i++ {
		// 不断计算能跳到的最远距离
		farthest = algorithm.Max(farthest, i+nums[i])
		// 可能碰到了 0，卡住跳不动了
		if farthest <= i {
			return false
		}
	}
	return farthest >= n-1
}
