/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := height(node.Left)
	right := height(node.Right)
	if left == -1 || right == -1 || left-right > 1 || left-right < -1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

// @lc code=end

