package linklist

// 剑指offer 22 . 返回链表的倒数第K个节点
func findFormEnd(head *ListNode,k int) *ListNode  {
	fast,slow := head,head
	for i := 0; i < k ;i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
