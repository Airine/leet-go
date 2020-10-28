package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindromeList(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }

    s, slow, fast := head, head, head.Next

    if fast.Next == nil {
        return slow.Val == fast.Val
    }

    for fast != nil && fast.Next != nil {
        s = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    middle := slow.Next
    if fast == nil {
        slow = s
    }
    slow.Next = nil
    tail := reverse(head)

    return compare(tail, middle)
}

func compare(a *ListNode, b *ListNode) bool {
    for a != nil && b != nil {
        if a.Val != b.Val {
            return false
        }
        a = a.Next
        b = b.Next
    }
    return true
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
    head := &ListNode{Next: &ListNode{Next: &ListNode{Next: nil, Val: 1}, Val: 0}, Val: 1}
    printListNode(head)
    isPalindromeList(head)
}
