package tree

//106. 从中序与后序遍历序列构造二叉树
func buildTreeFromInAndPost(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val:postorder[len(postorder)-1]}
	index := 0
	for i := range inorder {
		if inorder[i] == postorder[len(postorder)-1] {
			index = i
		}
	}

	root.Left = buildTreeFromInAndPost(inorder[0:index],postorder[0:index])
	root.Right = buildTreeFromInAndPost(inorder[index+1:],postorder[index:len(postorder)-1])
	return root
}
