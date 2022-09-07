package main

//左右根
func pruneTree(root *TreeNode) *TreeNode {
	if getChild(root) == 0 {
		return nil
	}
	return root
}

func getChild(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getChild(root.Left)
	right := getChild(root.Right)
	if left == 0 {
		root.Left = nil
	}
	if right == 0 {
		root.Right = nil
	}
	return root.Val + left + right
}
