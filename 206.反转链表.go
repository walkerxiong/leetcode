/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	node := head
	var last *ListNode
	var ret *ListNode
	for node != nil {
		nxt := node.Next
		node.Next = last
		last = node
		node = nxt
	}
	return last
}

// @lc code=end

