package tree

//572. 另一棵树的子树 https://leetcode.cn/problems/subtree-of-another-tree/
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var f func(root *TreeNode, subRoot *TreeNode) bool
	var check func(root *TreeNode, subRoot *TreeNode) bool
	check = func(root *TreeNode, subRoot *TreeNode) bool {
		if root == nil && subRoot == nil {
			return true
		} else if root == nil || subRoot == nil {
			return false
		}

		if root.Val != subRoot.Val {
			return false
		}

		left := check(root.Left, subRoot.Left)
		right := check(root.Right, subRoot.Right)

		return left && right
	}

	f = func(root *TreeNode, subRoot *TreeNode) bool {
		if root == nil {
			return false
		}

		if root.Val == subRoot.Val {
			if check(root, subRoot) {
				return true
			}
		}

		left := f(root.Left, subRoot)
		right := f(root.Right, subRoot)

		return left || right
	}
	return f(root, subRoot)
}
