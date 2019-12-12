package problem110

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
func isBalanced(root *TreeNode) bool {
	if root != nil {
		left := depth(root.Left)
		right := depth(root.Right)
		if left > right+1 || left < right-1 {
			return false
		}
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return true
}

func depth(root *TreeNode) int {
	if root != nil {
		left := depth(root.Left) + 1
		right := depth(root.Right) + 1
		if left > right {
			return left
		}
		return right
	}
	return 0
}
