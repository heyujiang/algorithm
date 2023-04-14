package linklist

//867. 链表的中间节点
func middleNode(head *ListNode)  *ListNode {
	fast,slow := head,head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
