package tree

//515. 在每个树行中找最大值 (https://leetcode.cn/problems/find-largest-value-in-each-tree-row/)
func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	var f func(root *TreeNode, depth int)
	f = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if len(res) <= depth {
			res = append(res, root.Val)
		} else if root.Val > res[depth] {
			res[depth] = root.Val
		}

		f(root.Left, depth+1)
		f(root.Right, depth+1)
	}

	f(root, 0)
	return res
}
