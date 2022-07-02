package linearstructs

import (
	"fmt"
	"testing"
)

/*
	A fila deve ser inicializada vazia
	Com a fila vazia e um novo elemento sendo inserido, ele deve ser o inicio e o fim da fila
	Com a fila preenchida e um novo elemento sendo inserido, ele deve ser o fim da fila
	Com a fila vazia, quando remoção, não deve ser possivel remover um elemento
	Com a fila preenchida com somente um elemento, quando removido, o fim e o inicio devem ser vazio
	Com a fila preenchida com mais de um elemento, quando houver uma remoção, o inicio deve apontar para o proximo elemento
*/
func TestInitQueue(t *testing.T) {
	q := InitQueue()

	if q == nil {
		t.Errorf("esperava uma fila existente")
	} else if q.First != nil {
		t.Errorf("esperava 'primeiro elemento vazio', obteve 'primeiro elemento %+v'", q)
	}
}

func TestInsertInEmptyQueue(t *testing.T){
	value := 10
	q := InitQueue()
	q.EnterQueue(value)

	if q.First.Value != value {
		expected := fmt.Sprintf("o primeiro membro da fila tivesse valor %d", value)
		obtained := fmt.Sprintf("primeiro membro da fila com valor %d", q.First.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} else if q.Last.Value != value {
		expected := fmt.Sprintf("o ultimo membro da fila tivesse valor %d", value)
		obtained := fmt.Sprintf("ultimo membro da fila com valor %d", q.Last.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	}
}

func TestInsertInFilledQueue(t *testing.T){
	previousValue, value := 10, 20
	q := InitQueue()
	q.EnterQueue(previousValue)
	q.EnterQueue(value)

	if q.First.Value != previousValue {
		expected := fmt.Sprintf("o primeiro membro da fila tivesse valor %d", previousValue)
		obtained := fmt.Sprintf("primeiro membro da fila com valor %d", q.First.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
	if q.Last.Value != value {
		expected := fmt.Sprintf("o ultimo membro da fila tivesse valor %d", value)
		obtained := fmt.Sprintf("ultimo membro da fila com valor %d", q.Last.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	}
	if q.First.Next != q.Last {
		expected := fmt.Sprintf("que o primeiro elemento, '%+v', fosse sucessido pelo ultimo elemento '%+v'", 
			q.First, q.Last)
		obtained := fmt.Sprintf("é sucessido pelo elemento '%+v'", q.First.Next)
		t.Errorf("esperava %s, mas %s", expected, obtained)

	}
}

func TestRemoveFromEmptyQueue(t *testing.T) {
	q := InitQueue()
	node := q.GetFromQueue()

	if node != nil {
		expected := fmt.Sprintf("não ser possivel remover de fila vazia")
		obtained := fmt.Sprintf("primeiro membro da fila sendo removido %+v", node)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	}
}

func TestRemoveFromQueueWithOneMember(t *testing.T) {
	value := 1
	q := InitQueue()
	q.EnterQueue(value)
	node := q.GetFromQueue()

	if node == nil {
		expected := fmt.Sprintf("existir um membro removido")
		obtained := fmt.Sprintf("membro removido é nulo")
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
	if node.Value != value {
		expected := fmt.Sprintf("membro removido com valor %d", value)
		obtained := fmt.Sprintf("membro removido tem valor %d", node.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
	if q.First != nil {
		expected := fmt.Sprintf("inicio da fila nulo")
		obtained := fmt.Sprintf("inicio da fila tem valor %+v", q.First)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
	if q.Last != nil {
		expected := fmt.Sprintf("fim da fila nulo")
		obtained := fmt.Sprintf("fim da fila tem valor %+v", q.First)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
}

func TestRemoveFromQueueWithTwoMembers(t *testing.T) {
	firstValue, secondValue := 1,2
	q := InitQueue()
	q.EnterQueue(firstValue)
	q.EnterQueue(secondValue)
	node := q.GetFromQueue()

	if node == nil {
		expected := fmt.Sprintf("existir um membro removido")
		obtained := fmt.Sprintf("membro removido é nulo")
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
	if node.Value != firstValue {
		expected := fmt.Sprintf("membro removido com valor %d", firstValue)
		obtained := fmt.Sprintf("membro removido tem valor %d", node.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
	if q.First.Value != secondValue {
		expected := fmt.Sprintf("inicio da fila com valor %d", secondValue)
		obtained := fmt.Sprintf("inicio da fila tem valor %d", q.First.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
	if q.Last.Value != secondValue {
		expected := fmt.Sprintf("fim da fila com valor %d", secondValue)
		obtained := fmt.Sprintf("fim da fila tem valor %d", q.First.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	}
}

func TestRemoveFromQueueWithNMembers(t *testing.T) {
	firstValue, secondValue, thirdValue := 1,2,3
	q := InitQueue()
	q.EnterQueue(firstValue)
	q.EnterQueue(secondValue)
	q.EnterQueue(thirdValue)
	node := q.GetFromQueue()

	if node == nil {
		expected := fmt.Sprintf("existir um membro removido")
		obtained := fmt.Sprintf("membro removido é nulo")
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
	if node.Value != firstValue {
		expected := fmt.Sprintf("membro removido com valor %d", firstValue)
		obtained := fmt.Sprintf("membro removido tem valor %d", node.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
	if q.First.Value != secondValue {
		expected := fmt.Sprintf("inicio da fila com valor %d", secondValue)
		obtained := fmt.Sprintf("inicio da fila tem valor %d", q.First.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	} 
	if q.Last.Value != thirdValue {
		expected := fmt.Sprintf("fim da fila com valor %d", thirdValue)
		obtained := fmt.Sprintf("fim da fila tem valor %d", q.First.Value)
		t.Errorf("esperava %s, mas %s", expected, obtained)
	}
}