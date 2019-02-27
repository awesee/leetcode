package maximum_depth_of_binary_tree

import . "github.com/openset/leetcode/internal/kit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left) + 1
	r := maxDepth(root.Right) + 1
	if l > r {
		return l
	}
	return r
}
