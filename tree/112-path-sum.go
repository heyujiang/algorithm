package tree

//112.路径总合（https://leetcode.cn/problems/path-sum/）
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var f func(root *TreeNode, sums int)
	var res bool
	f = func(root *TreeNode, sums int) {
		nodeSums := sums + root.Val
		if root.Left != nil && root.Right != nil {
			f(root.Left, nodeSums)
			f(root.Right, nodeSums)
		} else if root.Left != nil {
			f(root.Left, nodeSums)
		} else if root.Right != nil {
			f(root.Right, nodeSums)
		} else {
			if nodeSums == targetSum {
				res = true
			}
			return
		}
	}
	f(root, 0)
	return res
}

//113.路径总合II（https://leetcode.cn/problems/path-sum-ii/）
func hasPathSumII(root *TreeNode, targetSum int) [][]int {
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
