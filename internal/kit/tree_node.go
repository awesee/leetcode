package kit

import "math"

var NULL = math.MinInt32

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// convert []int to *TreeNode
func SliceInt2TreeNode(s []int) *TreeNode {
	l := len(s)
	if l == 0 {
		return nil
	}
	tree := &TreeNode{Val: s[0]}
	queue := make([]*TreeNode, 1)
	queue[0] = tree
	for i := 1; i < l; i++ {
		node := queue[0]
		queue = queue[1:]
		if s[i] != NULL {
			node.Left = &TreeNode{Val: s[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < l && s[i] != NULL {
			node.Right = &TreeNode{Val: s[i]}
			queue = append(queue, node.Right)
		}
	}
	return tree
}

// convert *TreeNode to []int
func TreeNode2SliceInt(t *TreeNode) (s []int) {
	return
}
