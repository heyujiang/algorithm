package linklist

// 19 && 剑指offer II 021.删除链表中倒数第n个节点
func removeNthFromEnd(head *ListNode,n int) *ListNode {
	dummy := &ListNode{Next:head}
	prev := findFormEnd(dummy,n+1)
	prev.Next = prev.Next.Next
	return dummy.Next
}