package tree

//543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	if root == nil {
		return 0
	}

	var maxDep func(node *TreeNode) int
	maxDep = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := maxDep(root.Left)
		rightDepth := maxDep(root.Right)

		if diameter < leftDepth + rightDepth {
			diameter = leftDepth + rightDepth
		}

		if leftDepth > rightDepth {
			return leftDepth + 1
		}
		return rightDepth + 1
	}

	maxDep(root)

	return diameter
}
