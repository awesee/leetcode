package problem437

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
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return findPath(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func findPath(root *TreeNode, sum int) int {
	r := 0
	if root == nil {
		return 0
	} else if root.Val == sum {
		r++
	}
	r += findPath(root.Left, sum-root.Val)
	r += findPath(root.Right, sum-root.Val)
	return r
}
