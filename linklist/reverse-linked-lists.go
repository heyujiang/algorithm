package linklist

// 206  && 剑指offer 24 && 剑指offer II 024
//反转链表

func reverseList(head *ListNode) *ListNode  {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList(head.Next)
	//1 2 3
	head.Next.Next = head
	head.Next = nil
	return last
}

func reverseListRange(head *ListNode) *ListNode {
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//temp,curr := new(ListNode),head
	//var prev *ListNode
	//for curr != nil {
	//	temp = curr.Next
	//	curr.Next = prev
	//	prev = curr
	//	curr = temp
	//}
	//return prev


	var prev *ListNode
	for head != nil {
		prev,head.Next,head = head,prev,head.Next
	}
	return prev
}

// 将一个链表前n个节点反转
var successor *ListNode
func reverseN(head *ListNode,n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}

	last := reverseN(head.Next,n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}
//1 2 3 4 5 6

// 92 反转链表II 将链表制定的区间反转
func reverseBetween(head *ListNode, left ,right int) *ListNode  {
	if left == 1 {
		return reverseN(head,right)
	}

	head.Next = reverseBetween(head.Next,left - 1 , right - 1)

	return head
}