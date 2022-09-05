package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getGoodNode(root.Left, root.Val) + getGoodNode(root.Right, root.Val) + 1
}

func getGoodNode(root *TreeNode, maxValue int) int {
	if root == nil {
		return 0
	}
	if root.Val >= maxValue {
		return getGoodNode(root.Left, root.Val) + getGoodNode(root.Right, root.Val) + 1
	} else {
		return getGoodNode(root.Left, maxValue) + getGoodNode(root.Right, maxValue)
	}
}
