package tree

import "math"

//222. 完全二叉树的节点个数

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh, rh := 0, 0
	lp := root
	for lp != nil {
		lh++
		lp = lp.Left
	}

	rp := root
	for rp != nil {
		rh++
		rp = rp.Right
	}

	if lh == rh {
		return int(math.Pow(2, float64(lh))) - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1

}
