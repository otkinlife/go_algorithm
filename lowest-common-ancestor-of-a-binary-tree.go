package go_algorithm

import "go_algorithm/common"

func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	leftRes := lowestCommonAncestor(root.Left, p, q)
	rightRes := lowestCommonAncestor(root.Right, p, q)
	if leftRes != nil && rightRes != nil {
		return root
	}
	if leftRes != nil {
		return leftRes
	}
	return rightRes
}
