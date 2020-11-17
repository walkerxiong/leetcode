package btree

type Stack struct {
	Data []*TreeNode
}

func (s *Stack) Pop() *TreeNode {
	val := s.Data[0]
	s.Data = s.Data[1:]
	return val
}

func (s *Stack) Push(v *TreeNode) {
	s.Data = append(s.Data, v)
}
