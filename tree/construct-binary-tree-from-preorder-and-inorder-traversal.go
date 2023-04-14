package tree

//105. 从前序与中序遍历序列构造二叉树   &&  剑指 Offer 07. 重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val:preorder[0]}

	index := 0
	for i := range inorder {
		if inorder[i] == preorder[0] {
			index = i
		}
	}

	root.Left = buildTree(preorder[1:index+1],inorder[0:index])
	root.Right = buildTree(preorder[index+1:],inorder[index+1:])
	return root
}
