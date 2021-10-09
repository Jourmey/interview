package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
