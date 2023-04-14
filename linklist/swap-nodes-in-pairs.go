package linklist
//24 . 两两交换链表中的节点
func swapPairs(head *ListNode)  *ListNode  {
	if head == nil || head.Next == nil{
		return head
	}
	a,b := head , head.Next
	nxt := b.Next
	b.Next = a
	a.Next = swapPairs(nxt)
	return b
}


