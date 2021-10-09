package main

//剑指 Offer 32 - I. 从上到下打印二叉树
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	tree(root, 0, func(n *TreeNode, level int) {
		for len(res) <= level {
			res = append(res, make([]int, 0))
		}
		res[level] = append(res[level], n.Val)
	})
	return res
}

func tree(node *TreeNode, level int, fu func(n *TreeNode, level int)) {
	if node == nil {
		return
	}
	fu(node, level)
	tree(node.Left, level+1, fu)
	tree(node.Right, level+1, fu)
}

func bfs(node *TreeNode, fu func(n *TreeNode, level int)) {
	if node == nil {
		return
	}
	que := make([]*TreeNode, 0, 16)
	que = append(que, node)
	for len(que) != 0 {
		o := que[0]
		que = que[1:]
		fu(o, len(que))

		if o.Left != nil {
			que = append(que, o.Left)
		}
		if o.Right != nil {
			que = append(que, o.Right)
		}
	}
}

func dfs(node *TreeNode, fu func(n *TreeNode, level int)) {
	if node == nil {
		return
	}
	que := make([]*TreeNode, 0, 16)
	que = append(que, node)
	for len(que) != 0 {
		o := que[len(que)-1]
		que = que[:len(que)-1]
		fu(o, len(que))

		if o.Right != nil {
			que = append(que, o.Right)
		}

		if o.Left != nil {
			que = append(que, o.Left)
		}
	}
}
