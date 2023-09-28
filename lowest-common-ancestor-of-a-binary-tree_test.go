package go_algorithm

import (
	"go_algorithm/common"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	case1 := common.ArrayToTreeNode([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	res1 := lowestCommonAncestor(case1, case1.Left, case1.Right)
	t.Log(res1.Val)
	case2 := common.ArrayToTreeNode([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	res2 := lowestCommonAncestor(case2, case2.Left, case2.Left.Right.Right)
	t.Log(res2.Val)
	case3 := common.ArrayToTreeNode([]interface{}{1, 2})
	res3 := lowestCommonAncestor(case3, case3, case3.Left)
	t.Log(res3.Val)
}
