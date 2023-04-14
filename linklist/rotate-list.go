package linklist

//61 . 循环链表
func rotateRight(head *ListNode,k int) *ListNode  {
	if head == nil {
		return head
	}

	l := 0
	a,tail := head,head
	for a != nil {
		l++
		tail = a
		a = a.Next
	}

	k = k%l
	if k == 0 {
		return head
	}

	kNode := findFormEnd(head,k+1)
	nHead := kNode.Next
	kNode.Next = nil
	tail.Next = head
	return nHead
}