package main


func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	nHead := next.Next
	next.Next = head
	head.Next = swapPairs(nHead)
	return next
}

func main() {
	
}
