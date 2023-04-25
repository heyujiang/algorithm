package tree

//404. 左叶子之和(https://leetcode.cn/problems/sum-of-left-leaves/)

func sumOfLeftLeaves(root *TreeNode) int {
	var res = 0
	var f func(root *TreeNode, isLeft bool)
	f = func(root *TreeNode, isLeft bool) {
		if root == nil {
			return
		}

		if isLeft && root.Left == nil && root.Right == nil {
			res += root.Val
			return
		}

		f(root.Left, true)
		f(root.Right, false)
	}

	f(root, false)
	return res
}
