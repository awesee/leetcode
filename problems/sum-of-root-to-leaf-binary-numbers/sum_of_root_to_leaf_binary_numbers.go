package problem1022

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
func sumRootToLeaf(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	sum := 0
	if root.Left != nil {
		root.Left.Val += root.Val << 1
		sum += sumRootToLeaf(root.Left)
	}
	if root.Right != nil {
		root.Right.Val += root.Val << 1
		sum += sumRootToLeaf(root.Right)
	}
	return sum
}
