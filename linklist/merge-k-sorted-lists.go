package linklist

import "container/heap"

//23 && 剑指offer II 078
//合并k个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{0,nil}
	head := dummy

	pq := make(PriorityQueue,0)
	heap.Init(&pq)
	for _,p := range lists {
		if p != nil {
			heap.Push(&pq,p)
		}
	}

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		head.Next = node
		if node.Next != nil {
			heap.Push(&pq,node.Next)
		}
		head = head.Next
	}
	
	return dummy.Next
}

type PriorityQueue []*ListNode

func (p PriorityQueue) Len() int{
	return len(p)
}

func (p PriorityQueue) Less(i,j int) bool  {
	return p[i].Val < p[j].Val
}

func (p PriorityQueue) Swap(i,j int)  {
	p[i],p[j] = p[j],p[i]
}

func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	node := old[n-1]
	*p = old[0:n-1]
	return node
}

func (p *PriorityQueue) Push(x interface{})  {
	node := x.(*ListNode)
	*p = append(*p,node)
}