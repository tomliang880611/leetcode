package leetcode

import (
	"fmt"
)

// this is done by referring to the solution and was
// adpated from python

// TreeNode is the structure for the node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	count := make(map[string]int)
	var ans []*TreeNode

	_, ans = collect(root, count, ans)
	return ans
}

func collect(node *TreeNode, count map[string]int, ans []*TreeNode) (string, []*TreeNode) {
	if node == nil {
		return "#", ans
	}

	leftVal, ans := collect(node.Left, count, ans)
	rightVal, ans := collect(node.Right, count, ans)
	serial := fmt.Sprintf("%v,%v,%v", node.Val, leftVal, rightVal)
	if val, ok := count[serial]; ok {
		if val == 1 {
			count[serial] = 2
			ans = append(ans, node)
		}
	} else {
		count[serial] = 1
	}

	return serial, ans
}
