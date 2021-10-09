package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Right, root.Left)
}

func symmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left != nil && right != nil {
		return left.Val == right.Val && symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
	}

	return false
}
