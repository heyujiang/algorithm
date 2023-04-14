package linklist

//142 && 剑指offer II 022. 环形链表II = 判断链表是否包含环，并返回环的起点，不包含返回null
func detectCycle(head *ListNode) *ListNode {

	fast,slow := head,head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			fast = head
			if head == slow {
				return slow
			}
			for {
				fast = fast.Next
				slow = slow.Next
				if fast == slow {
					return slow
				}
			}
		}
	}

	return nil
}

//141 . 环形链表I = 判断链表中是否包含环
func hasCycle(head *ListNode) bool {
	fast,slow := head,head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}