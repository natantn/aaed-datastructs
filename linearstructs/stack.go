package linearstructs

type StackNode struct {
	Value int
	Next  *StackNode
}

func CreateStackNode(value int) *StackNode {
	node := StackNode{value, nil}
	return &node
}

type Stack struct {
	Top            *StackNode
	LevelsQuantity int
}

func InitStack() *Stack {
	return &Stack{}
}

/*
	Inserção no Topo
	Remoção no Topo
*/

func (s *Stack) Push(value int) {
	node := CreateStackNode(value)
	node.Next = s.Top
	s.Top = node
	s.LevelsQuantity++
}

func (s *Stack) Pop() *StackNode {
	if s.Top == nil {
		return nil
	}

	node := s.Top
	s.Top = node.Next
	s.LevelsQuantity--
	return node
}
