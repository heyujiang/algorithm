package linklist

//21 && 剑指offer25 合并两个链表
func mergeTwoLists(node1 *ListNode,node2 *ListNode) *ListNode{
	dummy := &ListNode{0,nil}
	head := dummy
	for node1 != nil && node2 != nil {
		if node1.Val > node2.Val {
			head.Next = node2
			node2 = node2.Next
		}else {
			head.Next = node1
			node1 = node1.Next
		}
		head = head.Next
	}

	if node1 != nil {
		head.Next = node1
	}

	if node2 != nil {
		head.Next = node2
	}

	return dummy.Next
}