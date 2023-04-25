package tree

//501.二叉搜索数树中的众数（https://leetcode.cn/problems/find-mode-in-binary-search-tree/）
func findMode(root *TreeNode) []int {
	var base, count, maxCount int
	res := make([]int, 0)

	update := func(x int) {
		if x == base {
			count++
		} else {
			base = x
			count = 1
		}

		if count == maxCount {
			res = append(res, x)
		} else if count > maxCount {
			maxCount = count
			res = []int{base}
		}

	}

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		update(root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return res
}
