package btree

func preorderTraversal(root *TreeNode) []int {
	var s []*TreeNode
	node := root
	var l []int
	for node != nil || len(s) > 0 {
		for node != nil {
			l = append(l, node.Val)
			s = append(s, node)
			node = node.Left
		}
		node = s[len(s)-1].Right
		s = s[:len(s)-1]
	}
	return l
}

func inorderTraversal(root *TreeNode) []int {
	var l []int
	if root == nil {
		return l
	}
	left := root.Left
	if left != nil {
		inorder(left, &l)
	}
	l = append(l, root.Val)
	right := root.Right
	if right != nil {
		inorder(right, &l)
	}
	return l
}

func preorder(node *TreeNode, l *[]int) {
	v := node.Val
	*l = append(*l, v)
	if node.Left != nil {
		preorder(node.Left, l)
	}
	if node.Right != nil {
		preorder(node.Right, l)
	}
}

func inorder(node *TreeNode, l *[]int) {
	if node.Left != nil {
		inorder(node.Left, l)
	}

	v := node.Val
	*l = append(*l, v)

	if node.Right != nil {
		inorder(node.Right, l)
	}

}
