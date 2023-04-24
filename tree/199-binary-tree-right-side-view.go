package tree

//199. 二叉树的右视图(https://leetcode.cn/problems/binary-tree-right-side-view/)
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([]int, 0)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

			if i == n-1 {
				res = append(res, cur.Val)
			}
		}
	}
	return res
}
