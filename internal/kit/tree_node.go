package kit

import "math"

// NULL null
var NULL = math.MinInt32

// TreeNode - Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// SliceInt2TreeNode - convert []int to *TreeNode
func SliceInt2TreeNode(s []int) *TreeNode {
	l := len(s)
	if l == 0 {
		return nil
	}
	tree := &TreeNode{Val: s[0]}
	node, queue := tree, make([]*TreeNode, 0)
	for i := 1; i < l; i++ {
		if s[i] != NULL {
			node.Left = &TreeNode{Val: s[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < l && s[i] != NULL {
			node.Right = &TreeNode{Val: s[i]}
			queue = append(queue, node.Right)
		}
		if len(queue) > 0 {
			node = queue[0]
			queue = queue[1:]
		} else {
			break
		}
	}
	return tree
}

// TreeNode2SliceInt - convert *TreeNode to []int
func TreeNode2SliceInt(t *TreeNode) (s []int) {
	if t != nil {
		queue := []*TreeNode{t}
		for len(queue) > 0 {
			if queue[0] != nil {
				s = append(s, queue[0].Val)
				if queue[0].Left != nil || queue[0].Right != nil {
					queue = append(queue, queue[0].Left)
					queue = append(queue, queue[0].Right)
				}
			} else {
				s = append(s, NULL)
			}
			queue = queue[1:]
		}
	}
	for i := len(s) - 1; i >= 0 && s[i] == NULL; i-- {
		s = s[:i]
	}
	return
}
