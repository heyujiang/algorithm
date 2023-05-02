package tree

//653. 两数之和 IV - 输入二叉搜索树  https://leetcode.cn/problems/two-sum-iv-input-is-a-bst/
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if _, ok := m[k-root.Val]; ok {
			return true
		}
		m[root.Val] = struct{}{}

		return dfs(root.Left) || dfs(root.Right)
	}
	return dfs(root)
}
