package linearstructs

import "testing"

func TestCreateNode(t *testing.T) {
	value := 9
	node := CreateNode(9)

	if node.Value != value {
		t.Errorf("esperado '%d', obteve '%d'", value, node.Value)
	}
}

func TestInitList(t *testing.T) {
	linkedList := InitList()

	if linkedList == nil || linkedList.Head != nil {
		t.Errorf("esperado '%s', obteve '%s'", "Lista Ligada criada", "Valor Nulo")

	}
}

func TestInsertValueInEmptyList(t *testing.T) {
	l := InitList()
	l.InsertNodeWithValue(2)
	node, previousNode := l.SearchNodeByValue(2)

	if node == nil {
		t.Errorf("esperado '%s', obteve '%s'", "nó existente", "nó vazio")
	} else if node.Value != 2 && previousNode != nil {
		t.Errorf("esperado 'nó com valor %d', obteve 'nó com valor %d'", 2, node.Value)
	}
}
func TestInsertValueInListStart(t *testing.T) {
	l := InitList()
	l.InsertNodeWithValue(3)
	l.InsertNodeWithValue(2)
	node, previousNode := l.SearchNodeByValue(2)

	if node == nil {
		t.Errorf("esperado '%s', obteve '%s'", "nó existente", "nó vazio")
	} else if node.Value != 2 && previousNode != nil {
		t.Errorf("esperado 'nó com valor %d', obteve 'nó com valor %d'", 2, node.Value)
	} else if node.Prox.Value != 3 {
		t.Errorf("esperado 'nó proximo com valor %d', obteve 'nó com valor %d'", 3, node.Prox.Value)
	}
}

func TestInsertValueInListEnd(t *testing.T) {
	previousValue, value := 2,3
	l := InitList()
	l.InsertNodeWithValue(previousValue)
	l.InsertNodeWithValue(value)
	node, previousNode := l.SearchNodeByValue(value)

	if node == nil {
		t.Errorf("esperado '%s', obteve '%s'", "nó existente", "nó vazio")
	} else if node.Value != value {
		t.Errorf("esperado 'nó com valor %d', obteve 'nó com valor %d'", value, node.Value)
	} else if previousNode == nil || previousNode.Value != previousValue {
		t.Errorf("esperado 'nó anterior com valor %d', obteve 'nó com valor %d'", previousValue, previousNode.Value)
	}
}

func TestInsertValueInListMiddle(t *testing.T) {
	previousValue, value, nextValue := 2, 5, 10
	l := InitList()
	l.InsertNodeWithValue(previousValue)
	l.InsertNodeWithValue(nextValue)
	l.InsertNodeWithValue(value)
	node, previousNode := l.SearchNodeByValue(value)

	if node == nil {
		t.Errorf("esperado '%s', obteve '%s'", "nó existente", "nó vazio")
	} else if node.Value != value {
		t.Errorf("esperado 'nó com valor %d', obteve 'nó com valor %d'", value, node.Value)
	} else if previousNode == nil {
		t.Errorf("esperado '%s', obteve '%s'","nó anterior existente","nó anterior nulo")
		} else if previousNode.Value != previousValue {
		t.Errorf("esperado 'nó anterior com valor %d', obteve 'nó com valor %d'", previousValue, previousNode.Value)
	} else if node.Prox == nil {
		t.Errorf("esperado '%s', obteve '%s'","nó posterior existente","nó posterior nulo")
	} else if node.Prox.Value != nextValue {
		t.Errorf("esperado 'nó posterior com valor %d', obteve 'nó com valor %d'", nextValue, node.Prox.Value)
	}
}

func TestInsertDuplicatedValue(t *testing.T) {
	l := InitList()
	l.InsertNodeWithValue(3)
	if (l.InsertNodeWithValue(3)) {
		t.Errorf("inseriu nó duplicado")
	}
}

func TestRemoveValueInEmptyList(t *testing.T) {
	l := InitList()
	node := l.RemoveNodeByValue(0)

	if node != nil {
		t.Errorf("esperado '%s', obteve '%s'", "não ser possivel remover da lista vazia", "removeu algum nó")
	} 
}

func TestRemoveValueInListStart(t *testing.T) {
	l := InitList()
	l.InsertNodeWithValue(2)
	node, _ := l.SearchNodeByValue(2)
	removedNode := l.RemoveNodeByValue(2)

	if removedNode == nil {
		t.Errorf("esperado '%s', obteve '%s'", "nó existente", "nó vazio")
	} else if removedNode != node {
		t.Errorf("esperado 'nó com valor endereço %p e valor %d', obteve 'nó com valor endereço %p e valor %d",
		 &node, node.Value, &removedNode, removedNode.Value)
	} else if l.Head != nil {
		t.Errorf("esperado 'Lista vazia', obteve 'Lista apontando para o nó %+v'", l.Head)
	}
}

func TestRemoveValueInListEnd(t *testing.T) {
	previousValue, value := 2,3
	l := InitList()
	l.InsertNodeWithValue(previousValue)
	l.InsertNodeWithValue(value)
	node, previousNode := l.SearchNodeByValue(value)
	removedNode := l.RemoveNodeByValue(value)

	if removedNode == nil {
		t.Errorf("esperado '%s', obteve '%s'", "nó existente", "nó vazio")
	} else if removedNode != node {
		t.Errorf("esperado 'nó com valor endereço %p e valor %d', obteve 'nó com valor endereço %p e valor %d",
		 &node, node.Value, &removedNode, removedNode.Value)
	} else if previousNode.Prox != nil {
		t.Errorf("esperado 'nó anterior com prox vazio', obteve 'nó anterior com prox %+v'", previousNode.Prox)
	}
}

func TestRemoveValueInListMiddle(t *testing.T) {
	previousValue, value, nextValue := 2, 5, 10
	l := InitList()
	l.InsertNodeWithValue(previousValue)
	l.InsertNodeWithValue(nextValue)
	l.InsertNodeWithValue(value)
	node, previousNode := l.SearchNodeByValue(value)
	removedNode := l.RemoveNodeByValue(value)

	if removedNode == nil {
		t.Errorf("esperado '%s', obteve '%s'", "nó existente", "nó vazio")
	} else if removedNode != node {
		t.Errorf("esperado 'nó com valor endereço %p e valor %d', obteve 'nó com valor endereço %p e valor %d",
		 &node, node.Value, &removedNode, removedNode.Value)
	} else if previousNode == nil {
		t.Errorf("esperado '%s', obteve '%s'","nó anterior existente","nó anterior nulo")
	} else if previousNode.Prox != node.Prox {
		t.Errorf("esperado 'nó anterior com prox %+v', obteve 'nó anterior com prox %+v'", node.Prox, previousNode.Prox)
	} 
}

func TestPrintList(t *testing.T) {
	previousValue, value, nextValue := 2, 5, 10
	l := InitList()
	l.InsertNodeWithValue(previousValue)
	l.InsertNodeWithValue(nextValue)
	l.InsertNodeWithValue(value)

}
