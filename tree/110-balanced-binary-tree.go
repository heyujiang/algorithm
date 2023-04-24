package tree

import (
	"math"
)

//110. 平衡二叉树 (https://leetcode.cn/problems/balanced-binary-tree/)
func isBalanced(root *TreeNode) bool {
	var f func(root *TreeNode) (int, bool)
	f = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}
		left, isLeft := f(root.Left)
		right, isRight := f(root.Right)
		if isLeft && isRight {
			highDiff := int(math.Abs(float64(left - right)))
			if highDiff == 0 {
				return left + 1, true
			} else if highDiff == 1 {
				if left > right {
					return left + 1, true
				} else {
					return right + 1, true
				}
			} else {
				return 0, false
			}
		}
		return 0, false
	}
	_, res := f(root)
	return res
}
