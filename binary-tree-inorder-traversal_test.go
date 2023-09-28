package go_algorithm

import (
	"go_algorithm/common"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	case1 := common.ArrayToTreeNode([]interface{}{1, nil, 2, 3})
	res1 := InorderTraversal(case1)
	t.Log(res1)
	case2 := common.ArrayToTreeNode([]interface{}{})
	res2 := InorderTraversal(case2)
	t.Log(res2)
	case3 := common.ArrayToTreeNode([]interface{}{1})
	res3 := InorderTraversal(case3)
	t.Log(res3)
}
