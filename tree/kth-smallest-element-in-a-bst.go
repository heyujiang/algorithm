package tree

//230 二叉搜索数中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	sort := 0
	res := 0
	kthSmallestTraverse(root, k, &sort, &res)
	return res
}

func kthSmallestTraverse(root *TreeNode, k int, sort *int, res *int) {
	if root == nil {
		return
	}
	kthSmallestTraverse(root.Left, k, sort, res)
	*sort++
	if *sort == k {
		*res = root.Val
		return
	}

	kthSmallestTraverse(root.Right, k, sort, res)
}
