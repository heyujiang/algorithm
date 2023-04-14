package linklist
//
//func isPalindrome(head *ListNode) bool {
//	var f func(*ListNode) bool
//	left := head
//
//	f = func(right *ListNode) bool{
//		if right == nil {
//			return true
//		}
//
//		res := f(right.Next)
//		res = res && (right.Val == left.Val)
//		left = left.Next
//		return res
//	}
//	return f(head)
//}

var left *ListNode

func isPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(right *ListNode) bool  {
	if right == nil {
		return true
	}

	res := traverse(right.Next)
	res = res && (right.Val == left.Val)
	left = left.Next
	return res
}