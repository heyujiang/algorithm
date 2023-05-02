package tree

import "strconv"

//606.根据二叉树创建字符串  https://leetcode.cn/problems/construct-string-from-binary-tree/
func tree2str(root *TreeNode) string {
	var f func(root *TreeNode, isNil bool) string
	f = func(root *TreeNode, isNil bool) string {
		if root == nil {
			if isNil {
				return "()"
			}
			return ""
		}

		left, right := false, false
		if root.Right != nil && root.Left == nil {
			left = true
		}

		return "(" + strconv.Itoa(root.Val) + f(root.Left, left) + f(root.Right, right) + ")"
	}

	res := f(root, false)
	return res[1 : len(res)-1]
}
