package tree

import "math"

//124 . 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftPathSum := max(f(root.Left),0)     //如果路径和为负数，跟节点加上只会更小，因此返回0 ； 舍弃此自节点的路径和；
		rightPathSum := max(f(root.Right),0)

		res = max(res , root.Val + leftPathSum + rightPathSum)  // 路径和为自身的值加上左右子节点的路径和； 如果左右子节点的路径和为负数，则舍弃，计为零；
		return root.Val + max( leftPathSum , rightPathSum)  //左右两条路径，选择路径和大的那条返回；
	}

	f(root)

	return res
}

