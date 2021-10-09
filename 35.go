package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func main() {
	var a, b, c, d, e *Node
	e = &Node{
		Val:    1,
		Next:   nil,
		Random: nil,
	}
	d = &Node{
		Val:    10,
		Next:   e,
		Random: nil,
	}
	c = &Node{
		Val:    11,
		Next:   d,
		Random: nil,
	}
	b = &Node{
		Val:    13,
		Next:   c,
		Random: nil,
	}
	a = &Node{
		Val:    7,
		Next:   b,
		Random: nil,
	}

	aa := copyRandomList(a)
	fmt.Println(aa)
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{
			Val:    node.Val,
			Next:   node.Next,
			Random: nil,
		}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	n := head.Next
	for node := head; node != nil; node = node.Next {
		next := node.Next
		node.Next = next.Next
		if next.Next != nil {
			next.Next = next.Next.Next
		}
	}
	return n
}
