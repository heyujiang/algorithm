package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//116. 填充每个节点的下一个右侧节点指针
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	traverse(root.Left, root.Right)

	return root
}

func traverse(node1, node2 *Node) {
	if node1 == nil {
		return
	}
	node1.Next = node2

	traverse(node1.Left, node1.Right)
	traverse(node2.Left, node2.Right)
	traverse(node1.Right, node2.Left)
}

//117. 填充每个节点的下一个右侧节点指针 II (https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/)
func connectII(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)

		prev := queue[0]
		queue = queue[1:]
		if prev.Left != nil {
			queue = append(queue, prev.Left)
		}
		if prev.Right != nil {
			queue = append(queue, prev.Right)
		}

		for i := 1; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			prev.Next = cur
			prev = cur

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		prev.Next = nil
	}
	return root
}
