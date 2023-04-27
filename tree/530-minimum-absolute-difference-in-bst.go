package tree

import "math"

//530. 二叉搜索树的最小绝对差(https://leetcode.cn/problems/minimum-absolute-difference-in-bst)
func getMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt32
	inorder := make([]int, 0)
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		inorder = append(inorder, root.Val)
		traverse(root.Right)
	}
	traverse(root)

	for i := 0; i < len(inorder)-1; i++ {
		if inorder[i+1]-inorder[i] < minDiff {
			minDiff = inorder[i+1] - inorder[i]
		}
	}

	return minDiff
}
