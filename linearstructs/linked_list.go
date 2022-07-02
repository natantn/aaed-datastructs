package linearstructs

import "fmt"

type Node struct {
	Value int
	Prox  *Node
}

func CreateNode(value int) Node {
	node := Node{Value: value, Prox: nil}
	return node
}

type LinkedList struct {
	Head *Node
}

/*
	Inicialização OK
	Quantidade e exibição de elementos
	Busca OK
	Inserção OK
	Exclusão OK
*/

func InitList() (l *LinkedList) {
	l = &LinkedList{nil}
	return
}

func (l *LinkedList) PrintList() {
	actual := l.Head
	if actual == nil {
		fmt.Println("Lista vazia")
		return
	}
	for actual != nil {
		fmt.Println(actual)
	}
}

//
func (l *LinkedList) SearchNodeByValue(value int) (node, previousNode *Node) {
	actual := l.Head
	for actual != nil && actual.Value < value {
		previousNode = actual
		actual = actual.Prox
	}

	if actual != nil && actual.Value == value {
		node = actual
		return
	}

	node = nil
	return
}

func (l *LinkedList) InsertNodeWithValue(value int) bool {
	newNode := Node{value, nil}
	node, previousNode := l.SearchNodeByValue(newNode.Value)

	//Nó já existe e não é possivel inserir um novo
	if node != nil {
		return false
	}

	//Nó não existe
	if previousNode != nil { //Há um nó anterior para ele
		proxNode := previousNode.Prox
		previousNode.Prox = &newNode
		newNode.Prox = proxNode
		return true
	} else { //Será o primeiro nó da Lista
		proxNode := l.Head
		l.Head = &newNode
		newNode.Prox = proxNode
		return true
	}
}

func (l *LinkedList) RemoveNodeByValue(value int) (deletedNode *Node) {
	node, previousNode := l.SearchNodeByValue(value)

	if node == nil { //Nó não foi encontrado
		return nil
	}

	if previousNode == nil { //Não há nó anterior, logo o nó encontrado é a cabeça da lista
		l.Head = nil
	} else {
		previousNode.Prox = node.Prox
	}
	return node
}
