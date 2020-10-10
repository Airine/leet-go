package main

// Definition for singly-linked list.
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow = slow.Next
	another := head
	for slow != another {
		slow = slow.Next
		another = another.Next
	}
	return another
}

func main() {
	
}
