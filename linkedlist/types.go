package linkedlist

import "cmp"

type Node[T cmp.Ordered] struct {
	value             T
	prvNode, nextNode *Node[T]
}

type SLinkedList[T cmp.Ordered] struct {
	size               uint16
	headNode, tailNode *Node[T]
}

type LinkedList[T cmp.Ordered] interface {
	AddAtFirst(value T)
	AddAtLast(value T)
	DeleteAtFirst()
	DeleteAtLast()
	NodeContains(value T)
	isEmpty()
	PrintSize()
}

func CreateNode[T cmp.Ordered](value T) *Node[T] {
	return &Node[T]{value: value, prvNode: nil, nextNode: nil}
}
