package tree

//129. 求根节点到叶节点数字之和(https://leetcode.cn/problems/sum-root-to-leaf-numbers/)
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	var f func(root *TreeNode, sums int)
	f = func(root *TreeNode, sums int) {
		sums = sums*10 + root.Val

		if root.Left != nil && root.Right != nil {
			f(root.Left, sums)
			f(root.Right, sums)
		} else if root.Left != nil {
			f(root.Left, sums)
		} else if root.Right != nil {
			f(root.Right, sums)
		} else {
			res += sums
		}
	}

	f(root, 0)
	return res
}
