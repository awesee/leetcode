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
	queue := make([]*TreeNode, 0)
	node := tree
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

// convert *TreeNode to []int
func TreeNode2SliceInt(t *TreeNode) (s []int) {
	if t != nil {
		queue := make([]*TreeNode, 1)
		queue[0] = t
		for len(queue) > 0 {
			if queue[0] != nil {
				s = append(s, queue[0].Val)
				if queue[0].Left != nil || queue[0].Right != nil {
					queue = append(queue, queue[0].Left)
				}
				if queue[0].Right != nil {
					queue = append(queue, queue[0].Right)
				}
			} else {
				s = append(s, NULL)
			}
			queue = queue[1:]
		}
	}
	return
}
