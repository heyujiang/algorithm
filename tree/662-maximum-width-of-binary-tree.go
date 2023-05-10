package tree

//662. 二叉树最大宽度  https://leetcode.cn/problems/maximum-width-of-binary-tree/

//解法一：广度优先+节点编号
func widthOfBinaryTreeB(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	root.Val = 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		width := queue[n-1].Val - queue[0].Val + 1
		if res < width {
			res = width
		}
		for i := 0; i < n; i++ {
			cur := queue[i]
			if cur.Left != nil {
				cur.Left.Val = 2 * cur.Val
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				cur.Right.Val = 2*cur.Val + 1
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[n:]
	}
	return res
}

//解法一：深度优先+节点编号
func widthOfBinaryTreeD(root *TreeNode) int {
	if root == nil {
		return 0
	}
	imap := make(map[int]int)
	res := 1
	var traverse func(root *TreeNode, index int, depth int)
	traverse = func(root *TreeNode, index int, depth int) {
		if root == nil {
			return
		}

		if _, ok := imap[depth]; ok {
			width := index - imap[depth] + 1
			if res < width {
				res = width
			}
		} else {
			imap[depth] = index
		}

		traverse(root.Left, index*2, depth+1)
		traverse(root.Right, index*2+1, depth+1)
	}

	traverse(root, 1, 0)
	return res
}
