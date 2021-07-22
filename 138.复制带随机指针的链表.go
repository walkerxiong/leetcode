/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	var m = make(map[*Node]*Node)
	currNode := head
	for currNode != nil {
		m[currNode] = &Node{Val: currNode.Val, Next: currNode.Next, Random: currNode.Random}
		currNode = currNode.Next
	}
	curr := m[head]
	for curr != nil {
		curr.Next = m[curr.Next]
		curr.Random = m[curr.Random]
		curr = curr.Next
	}
	return m[head]
}

// @lc code=end

