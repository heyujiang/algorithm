package tree

//133 . 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{root.Val}
	res = append(res,preorderTraversal(root.Left)...)
	res = append(res,preorderTraversal(root.Right)...)
	return res
}

/*
func traversal(root *TreeNode)  {
	if root == nil {
		return
	}

	//前序遍历
	traversal(root.Left)
	//中序遍历
	traversal(root.Right)
	//后序遍历
}
*/