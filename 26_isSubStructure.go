package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//剑指 Offer 26. 树的子结构
//输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
//
//B是A的子结构， 即 A中有出现和B相同的结构和节点值。

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	if isSub(A, B) {
		return true
	}

	return isSubStructure(A.Right, B) || isSubStructure(A.Left, B)
}

func isSub(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}

	return a.Val == b.Val && isSub(a.Left, b.Left) && isSub(a.Right, b.Right)
}
