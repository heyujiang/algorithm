package tree

//113.路径总和II（https://leetcode.cn/problems/path-sum-ii/）
func pathSumII(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	var f func(root *TreeNode, sums int, path []int)
	var res [][]int
	f = func(root *TreeNode, sums int, path []int) {
		nodeSums := sums + root.Val
		path = append(path, root.Val)
		if root.Left != nil && root.Right != nil {
			f(root.Left, nodeSums, path)
			f(root.Right, nodeSums, path)
		} else if root.Left != nil {
			f(root.Left, nodeSums, path)
		} else if root.Right != nil {
			f(root.Right, nodeSums, path)
		} else {
			if nodeSums == targetSum {
				res = append(res, append([]int{}, path...))
			}
			return
		}
	}
	f(root, 0, []int{})
	return res
}
