package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}
