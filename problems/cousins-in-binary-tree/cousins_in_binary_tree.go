package problem993

import "github.com/openset/leetcode/internal/kit"

// TreeNode - Definition for a binary tree node.
type TreeNode = kit.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	xp, yp, xd, yd, depth := 0, 0, 0, 0, 0
	queue := []*TreeNode{root}
	for l := len(queue); l > 0 && xd == 0 && yd == 0; l = len(queue) {
		for _, node := range queue {
			if node.Left != nil {
				if node.Left.Val == x {
					xp, xd = node.Val, depth+1
				} else if node.Left.Val == y {
					yp, yd = node.Val, depth+1
				} else {
					queue = append(queue, node.Left)
				}
			}
			if node.Right != nil {
				if node.Right.Val == x {
					xp, xd = node.Val, depth+1
				} else if node.Right.Val == y {
					yp, yd = node.Val, depth+1
				} else {
					queue = append(queue, node.Right)
				}
			}
		}
		depth, queue = depth+1, queue[l:]
	}
	return xp != yp && xd == yd
}
