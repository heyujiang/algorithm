package tree

//563.二叉树的坡度   https://leetcode.cn/problems/binary-tree-tilt/
func findTilt(root *TreeNode) int {
	res := 0
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := f(root.Left)
		right := f(root.Right)
		if right > left {
			res += right - left
		} else {
			res += left - right
		}

		return root.Val + left + right
	}
	f(root)
	return res
}
