package linklist

import (
	"math"
)

// 82 . 删除排序两种重复元素 II
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return  head
	}
	dummy := &ListNode{math.MinInt32,head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		}else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 83 . 删除排序联保中的重复元素
func deleteDuplicates(head *ListNode) *ListNode  {
	if head == nil {
		return nil
	}
	fast,slow := head,head

	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}