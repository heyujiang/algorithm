package linklist

//160 && 剑指offer 52 && 剑指offer II 023
//相交链表

//a1 a2 c1 c2 c3 b1 b2 b3 c1 v2 c3
//b1 b2 b3 c1 v2 c3 a1 a2 c1 c2 c3
func getIntersectionNode(headA , headB *ListNode ) *ListNode  {
	a,b := headA,headB
	for a != b {
		if a == nil {
			a = headB
		}else {
			a = a.Next
		}

		if b == nil {
			b = headA
		}else {
			b = b.Next

		}
	}
	return a
}