package main

import (
	"fmt"
)

//
type ListNode struct {
	Val int
	Next *ListNode
}

func merge(head *ListNode, tail *ListNode)  {
	h, t := head, tail
	for h != nil && t != nil {
		hn := h.Next
		tn := t.Next
		h.Next = t
		if hn != nil {
			t.Next = hn
		}
		h = hn
		t = tn
	}
}

func reverse(head *ListNode) *ListNode  {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	head.Next = nil
	for second != nil {
		s := second.Next
		second.Next = first
		first = second
		second = s
	}
	return first
}

func reorderList(head *ListNode)  {
	l := 0
	target := head
	for target != nil {
		target = target.Next
		l ++
	}
	if l <= 2 {
		return
	}
	target = head
	for i := 0; i < l/2 - 1; i++ {
		target = target.Next
	}
	rev := target.Next
	target.Next = nil
	tail := reverse(rev)
	merge(head, tail)
}

func printListNode(head *ListNode)  {
	fmt.Print(head.Val)
	target := head.Next
	for target != nil {
		fmt.Printf("-> %d", target.Val)
		target = target.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{Next: &ListNode{Next: &ListNode{Next: &ListNode{Next: &ListNode{Next: nil, Val: 5}, Val: 4}, Val: 3}, Val: 2}, Val: 1}
	printListNode(head)
	reorderList(head)
	printListNode(head)
}
