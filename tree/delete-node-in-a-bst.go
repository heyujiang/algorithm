package tree

//450  删除二叉树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Left
		} else if root.Right == nil {
			return root.Right
		} else {
			maxNode := root.Left
			for maxNode.Right != nil {
				maxNode = maxNode.Right
			}

			root.Left = deleteNode(root.Left, maxNode.Val)
			maxNode.Left = root.Left
			maxNode.Right = root.Right
			root = maxNode
		}
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
