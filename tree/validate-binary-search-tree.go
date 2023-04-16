package tree

//98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return validBST(root, nil, nil)
}

func validBST(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	} else if max != nil && root.Val >= max.Val {
		return false
	}

	return validBST(root.Left, min, root) && validBST(root.Right, root, max)
}
