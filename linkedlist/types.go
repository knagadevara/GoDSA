package linkedlist

type IntNode struct {
	value             int64
	prvNode, nextNode *IntNode
}

type SLinkedList struct {
	headNode, tailNode *IntNode
}

type LinkedList interface {
	setInitialNode(node *IntNode)
	AddAtFirst(value int64)
	DeleteAtFirst()
	AddAtLast(value int64)
	DeleteAtLast()
	NodeContains(value int64)
	PrintValues()
}
