package linkedlist

type Node[T comparable] struct {
	value             T
	prvNode, nextNode *Node[T]
}

type SLinkedList[T comparable] struct {
	size               uint16
	headNode, tailNode *Node[T]
}

type LinkedList[T comparable] interface {
	AddAtFirst(value T)
	AddAtLast(value T)
	DeleteAtFirst()
	DeleteAtLast()
	NodeContains(value T)
	isEmpty()
	PrintSize()
}

func CreateNode[T comparable](value T) *Node[T] {
	return &Node[T]{value: value, prvNode: nil, nextNode: nil}
}
