package linkedlist

type IntNode struct {
	value             int64
	prvNode, nextNode *IntNode
}

type SLinkedList struct {
	totalNodes         int16
	headNode, tailNode *IntNode
}

type LinkedList interface {
	CreateNewNode(value int64) *IntNode
	AddAtFirst(value int64)
	DeleteAtFirst()
	AddAtLast(value int64)
	DeleteAtLast()
	NodeContains(value int64)
	IndexOf(value int64)
}
