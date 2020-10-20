package main

import (
    "fmt"
)

//
type ListNode struct {
    Val int
    Next *ListNode
}

func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    // 快慢指针找到中点
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }
    // 从中点断开
    middle := slow.Next
    slow.Next = nil
    // reverse（逆转）中点以后的list
    tail := reverse(middle)
    // merge（归并）两个list
    merge(head, tail)
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
    head := &ListNode{Next: &ListNode{Next: &ListNode{Next: &ListNode{Next: nil, Val: 4}, Val: 3}, Val: 2}, Val: 1}
    printListNode(head)
    reorderList(head)
    printListNode(head)
}
