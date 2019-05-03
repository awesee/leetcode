package range_sum_of_bst

import . "github.com/openset/leetcode/internal/kit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	ans := 0
	if root.Val >= L && root.Val <= R {
		ans += root.Val
	}
	if root.Left != nil && root.Val > L {
		ans += rangeSumBST(root.Left, L, R)
	}
	if root.Right != nil && root.Val < R {
		ans += rangeSumBST(root.Right, L, R)
	}
	return ans
}
