package tree

//889. 根据前序和后序遍历构造二叉树
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val:preorder[0]}
	if len(preorder) == 1  {
		return root
	}
	val := preorder[1]
	index :=0
	for i := range postorder {
		if val == postorder[i] {
			index = i
		}
	}


	root.Left = constructFromPrePost(preorder[1:index+2],postorder[0:index+1])
	root.Right = constructFromPrePost(preorder[index+2:],postorder[index+1:len(postorder)-1])
	return root
}