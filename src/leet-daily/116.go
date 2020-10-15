package main

type node struct {
	Val int
	Left *node
	Right *node
	Next *node
}

func connect(root *node) *node {
	if root == nil || root.Left == nil {
		return root
	}
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	} else {
		root.Right.Next = nil
	}
	root.Left.Next = connect(root.Right)
	root.Left = connect(root.Left)
	return root
}

func main() {
	
}
