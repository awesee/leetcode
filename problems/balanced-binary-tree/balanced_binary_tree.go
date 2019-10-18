package problem_110

import "github.com/openset/leetcode/internal/kit"

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
	_, isBalanced := recur(root)
	return isBalanced
}

func recur(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftIsBalanced := recur(root.Left)
	rightDepth, rightIsBalanced := recur(root.Right)
	if leftIsBalanced && rightIsBalanced &&
		-1 <= leftDepth-rightDepth && leftDepth-rightDepth <= 1 {
		return max(leftDepth, rightDepth) + 1, true
	}
	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
