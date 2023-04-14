package linklist

//25
// K个一组反转链表

func reverse(head *ListNode) *ListNode {
	var pre,cur,nxt *ListNode
	pre,cur,nxt = nil,head,nil
	for cur != nil {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reverseAB(a , b *ListNode) *ListNode  {
	var pre,cur,nxt *ListNode
	pre,cur,nxt = nil,a,nil
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func reverseKGroup(head *ListNode, k int) *ListNode{
	a,b := head,head
	for i:=0;i<k;i++{
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverseAB(a,b)
	a.Next = reverseKGroup(b,k)
	return newHead
}