package btree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var l []int
	preoder(root, &l)
	return l
}

func preoder(node *TreeNode, l *[]int) {
	v := node.Val
	*l = append(*l, v)
	if node.Left != nil {
		preoder(node.Left, l)
	}
	if node.Right != nil {
		preoder(node.Right, l)
	}
}
