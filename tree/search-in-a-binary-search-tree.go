package tree

//700. 二叉搜索树中的搜索
func searchTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < val {
		return searchTree(root.Right, val)
	} else if root.Val > val {
		return searchTree(root.Left, val)
	}

	return root
}
