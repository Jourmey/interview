package main

func reversePrint(head *ListNode) []int {
	n := reverseList(head)
	i := make([]int, 0, 16)

	for n != nil {
		i = append(i, n.Val)
		n = n.Next

	}

	return i
}
