package main

func mirrorTree(root *TreeNode) *TreeNode {
	tree(root, 0, func(n *TreeNode, level int) {
		n.Right, n.Left = n.Left, n.Right
	})
	return root
}
