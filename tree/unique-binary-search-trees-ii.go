package tree

//95. 不同的二叉搜索树 II
func generateTrees(n int) []*TreeNode {
	return buildBST(1, n)
}

func buildBST(lo, hi int) []*TreeNode {
	var res []*TreeNode
	if lo > hi {
		res = append(res, nil)
		return res
	}

	for i := lo; i <= hi; i++ {
		leftMaps := buildBST(lo, i-1)
		rightMaps := buildBST(i+1, hi)
		for _, left := range leftMaps {
			for _, right := range rightMaps {
				res = append(res, &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}
	return res
}
