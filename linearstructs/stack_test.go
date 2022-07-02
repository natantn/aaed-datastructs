package linearstructs

import "testing"

func TestInitStack(t *testing.T) {
	s := InitStack()

	if s == nil {
		t.Errorf("esperava '%s'", "uma pilha existente")
	} else if s.Top != nil {
		t.Errorf("esperava '%s', obtido 'uma pilha com topo '%+v'", "uma pilha vazia", s.Top)
	}
}

func TestPushEmptyStack(t *testing.T){
	value := 1
	s := InitStack()
	s.Push(value)

	if s.Top.Value != value {
		t.Errorf("esperava 'valor %d no topo', obteve 'valor %d no topo'", value, s.Top.Value)
	}
}

func TestPushFilledStack(t *testing.T) {
	previousValue, value := 1,2
	s := InitStack()
	s.Push(previousValue)
	s.Push(value)

	if s.Top.Value != value {
		t.Errorf("esperava 'valor %d no topo', obteve 'valor %d no topo'", value, s.Top.Value)
	} else if s.Top.Next.Value != previousValue {
		t.Errorf("esperava 'valor %d no nivel inferior ao topo', obteve 'valor %d no nivel inferior ao topo'",
			previousValue, s.Top.Next.Value)
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := InitStack()
	node := s.Pop()

	if node != nil {
		t.Errorf("esperava 'não ser possivel remover do topo', obteve topo com valor '%+v'", node)
	}
}

func TestPopStackFilledWithOneLevel(t *testing.T){
	s := InitStack()
	s.Push(1)
	stackTop := s.Top

	node := s.Pop()

	if node == nil || node != stackTop {
		t.Errorf("esperava 'um nó de topo', obteve '%+v''", node)
	} else if node.Value != 1 {
		t.Errorf("esperava 'nó com valor %d', obteve 'nó com valor %d'", 1, node.Value)
	} else if s.Top != nil {
		t.Errorf("esperava 'pilha vazia', obteve 'pilha com topo %+v'", s.Top)
	}
}

func TestPopStackFilledWithNLevel(t *testing.T){
	previousValue, value := 1,2
	s := InitStack()
	s.Push(previousValue)
	s.Push(value)
	stackTop := s.Top

	node := s.Pop()

	if node == nil || node != stackTop {
		t.Errorf("esperava 'um nó de topo', obteve '%+v''", node)
	} else if node.Value != value {
		t.Errorf("esperava 'nó com valor %d', obteve 'nó com valor %d'", value, node.Value)
	} else if s.Top.Value != previousValue {
		t.Errorf("esperava 'pilha com topo %d', obteve 'pilha com topo %d'", previousValue, s.Top.Value)
	}
}
