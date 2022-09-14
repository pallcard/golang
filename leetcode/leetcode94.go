package main

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	left := inorderTraversal(root.Left)
	res = append(res, left...)
	res = append(res, root.Val)
	right := inorderTraversal(root.Right)
	res = append(res, right...)
	return res
}
