package tree

import "fmt"

//257.二叉树路径路径 https://www.leetcode.cn/problems/binary-tree-paths
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)

	if root == nil {
		return res
	}

	var traverse func(root *TreeNode, str string)
	traverse = func(root *TreeNode, str string) {
		if root == nil {
			return
		}

		if str == "" {
			str = fmt.Sprintf("%d", root.Val)
		} else {
			str += fmt.Sprintf("->%d", root.Val)
		}

		if root.Left == nil && root.Right == nil {
			res = append(res, str)
			return
		}

		traverse(root.Left, str)
		traverse(root.Right, str)
	}

	traverse(root, "")
	return res
}
