package tree

//538. 把二叉搜索树转换为累加树  &&  1038. 从二叉搜索树到更大和树 &&   剑指 Offer II 054. 所有大于等于节点的值之和
var tempN int

func convertBST(root *TreeNode) *TreeNode {
	tempN = 0
	convertBSTTraverse(root)
	return root
}

func convertBSTTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	convertBSTTraverse(root.Right)
	root.Val += tempN
	tempN = root.Val
	convertBSTTraverse(root.Left)
}
