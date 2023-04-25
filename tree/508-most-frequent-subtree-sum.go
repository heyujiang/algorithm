package tree

//508. 出现次数最多的子树元素和(https://leetcode.cn/problems/most-frequent-subtree-sum/)
func findFrequentTreeSum(root *TreeNode) []int {
	counts := make(map[int]int, 0)
	maxCount := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)

		sums := left + right + root.Val

		counts[sums]++
		if counts[sums] > maxCount {
			maxCount = counts[sums]
		}
		return sums
	}

	dfs(root)
	var res []int
	for k, v := range counts {
		if v == maxCount {
			res = append(res, k)
		}
	}
	return res

}
