package go_algorithm

import (
	"go_algorithm/common"
)

// InorderTraversal 二叉树的中序遍历
// 中序遍历（Inorder Traversal）是一种树的遍历方式，主要应用于二叉树。
// 在中序遍历中，我们按照 左子树-根节点-右子树 的顺序访问二叉树的节点。具体步骤如下：
//  1. 遍历左子树
//  2. 访问根节点
//  3. 遍历右子树
func InorderTraversal(root *common.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	res = append(InorderTraversal(root.Left), root.Val)
	res = append(res, InorderTraversal(root.Right)...)
	return res
}
