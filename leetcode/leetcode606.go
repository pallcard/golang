package main

import "fmt"

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	left := tree2str(root.Left)
	right := tree2str(root.Right)
	if left != "" {
		left = fmt.Sprintf("(%s)", left)
	}
	if right != "" {
		right = fmt.Sprintf("(%s)", right)
	}
	return fmt.Sprintf("%d%s%s", root.Val, left, right)
}
