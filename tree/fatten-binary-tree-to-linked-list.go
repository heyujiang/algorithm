package tree

func flatten1(root *TreeNode) {
	var f func(node *TreeNode) (*TreeNode,*TreeNode)
	f = func(node *TreeNode) (*TreeNode, *TreeNode) {
		if node == nil {
			return nil,nil
		}

		leftHead,leftTail := f(node.Left)
		rightHead,rightTail := f(node.Right)

		if leftHead == nil && rightHead == nil && leftTail == nil && rightTail == nil {
			return node,node
		}else if leftHead == nil && leftTail == nil {
			return node,rightTail
		}

		node.Left = nil
		node.Right = leftHead
		leftTail.Right = rightHead

		if rightTail == nil {
			return node,leftTail
		}
		return node,rightTail
	}

	f(root)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

}