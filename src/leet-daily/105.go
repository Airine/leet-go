package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}
	val := preorder[0]
	pos := 0
	for i := 0; i < len(inorder); i ++ {
		if val == inorder[i] {
			pos = i
			break
		}
	}
	return &TreeNode{
		val,
		buildTree(preorder[1:pos+1], inorder[:pos+1]),
		buildTree(preorder[pos+1:], inorder[pos+1:])}
}
func printTree(node *TreeNode) string {
	return "\nvalue: " + string(node.Val) +
		",left: (" + printTree(node.Left) + ")" +
		",right: (" + printTree(node.Right) + ")\n"
}

func main() {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}
	root := buildTree(preorder, inorder)
	fmt.Println(root)
}
