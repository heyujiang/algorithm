package tree

//513. 找树左下角的值(https://leetcode.cn/problems/find-bottom-left-tree-value/)

func findBottomLeftValue(root *TreeNode) int {
	var res, maxDepth int
	var f func(root *TreeNode, depth int)
	f = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		f(root.Left, depth+1)
		if depth > maxDepth {
			res = root.Val
			maxDepth = depth
		}
		f(root.Right, depth+1)
	}

	f(root, 1)
	return res
}
