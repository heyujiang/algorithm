package linklist

//86. 分割链表
func partition(head *ListNode,x int) *ListNode {
	dummy,tail := &ListNode{},&ListNode{}
	d,t,p := dummy,tail,head
	for p != nil  {
		if p.Val < x {
			d.Next = p
			d = d.Next
		}else{
			t.Next = p
			t = t.Next
		}
		temp := p.Next
		p.Next = nil
		p = temp
	}

	d.Next = tail.Next

	return dummy.Next
}