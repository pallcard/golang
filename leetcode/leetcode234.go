package main

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	arr = append(arr, head.Val)
	p := head
	for p.Next != nil {
		arr = append(arr, p.Next.Val)
		p = p.Next
	}

	i := len(arr)/2 - 1
	j := len(arr) / 2
	if len(arr)%2 != 0 {
		i = len(arr) / 2
	}

	for i >= 0 {
		if arr[i] != arr[j] {
			return false
		}
		i--
		j++
	}
	return true
}
