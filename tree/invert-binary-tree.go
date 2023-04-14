package tree

// 226 . 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	//invertTraverse(root)
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)

	root.Left,root.Right = root.Right,root.Left
	return root
}

func invertTraverse(root *TreeNode){
	if root == nil {
		return
	}

	//root.Left,root.Right = root.Right,root.Left

	invertTree(root.Left)
	//right := root.Right
	//root.Left,root.Right = root.Right,root.Left
	//invertTree(right)
	invertTree(root.Right)

	root.Left,root.Right = root.Right,root.Left
}