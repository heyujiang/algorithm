package tree

//623.在二叉树中增加一行 https://leetcode.cn/problems/add-one-row-to-tree/
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}

	var f func(root *TreeNode, treeDep int)
	f = func(root *TreeNode, treeDep int) {
		if root == nil {
			return
		}
		if treeDep == depth-1 {
			root.Left = &TreeNode{Val: val, Left: root.Left}
			root.Right = &TreeNode{Val: val, Right: root.Right}
			return
		}

		f(root.Left, treeDep+1)
		f(root.Right, treeDep+1)
	}
	f(root, 1)
	return root
}
