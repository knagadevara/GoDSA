package linkedlist

import "fmt"

func (n *IntNode) setValue(value int64) {
	n.value = value
}

func (n *IntNode) getValue() int64 {
	return n.value
}

func (n *IntNode) setNextNode(nextNode *IntNode) {
	n.nextNode = nextNode
}

func (n *IntNode) setPrvNode(prvNode *IntNode) {
	n.prvNode = prvNode
}

func (n *IntNode) getNextNode() *IntNode {
	return n.nextNode
}

func (n *IntNode) getPrvNode() *IntNode {
	return n.prvNode
}

func (n *SLinkedList) setHeadNode(headNode *IntNode) {
	n.headNode = headNode
}

func (n *SLinkedList) setTailNode(tailNode *IntNode) {
	n.tailNode = tailNode
}

func (n *SLinkedList) getHeadNode() *IntNode {
	return n.headNode
}

func (n *SLinkedList) getTailNode() *IntNode {
	return n.tailNode
}

func CreateNode(value int64) *IntNode {
	return &IntNode{value: value, prvNode: nil, nextNode: nil}
}

func (n *SLinkedList) setInitialNode(node *IntNode) {
	node.setPrvNode(node)
	node.setNextNode(node)
	n.setHeadNode(node)
	n.setTailNode(node)
}

func (n *SLinkedList) AddAtFirst(value int64) {
	node := CreateNode(value)
	if n.isEmpty() {
		n.setInitialNode(node)
	} else {
		node.setNextNode(n.getHeadNode())
		n.getHeadNode().setPrvNode(node)
		n.setHeadNode(node)
	}
	fmt.Println("Added: ", value)
}

func (n *SLinkedList) AddAtLast(value int64) {
	node := CreateNode(value)
	if n.isEmpty() {
		n.setInitialNode(node)
	} else {
		node.setPrvNode(n.getTailNode())
		n.getTailNode().setNextNode(node)
		n.setTailNode(node)
	}
	fmt.Println("Added: ", value)
	// n.IncNode()
}

func (n *SLinkedList) isEmpty() bool {
	return n.getHeadNode() == nil || n.getTailNode() == nil
}

func (n *SLinkedList) DeleteAtFirst() {
	if !(n.isEmpty()) {
		fmt.Printf("Deleting: %d\n", n.getHeadNode().getValue())
		secondNode := n.getHeadNode().getNextNode()
		n.getHeadNode().setPrvNode(nil)
		n.getHeadNode().setValue(0)
		n.getHeadNode().setNextNode(nil)
		n.setHeadNode(secondNode)
	} else {
		fmt.Printf("Noting to Delete!!\n")
	}
}

func (n *SLinkedList) DeleteAtLast() {
	if !(n.isEmpty()) {
		fmt.Printf("Deleting: %d\n", n.getTailNode().getValue())
		lastButOne := n.getTailNode().getPrvNode()
		n.getTailNode().setNextNode(nil)
		n.getTailNode().setValue(0)
		n.getTailNode().setNextNode(nil)
		n.setTailNode(lastButOne)
	} else {
		fmt.Printf("Noting to Delete!!\n")
	}
}

func (n *SLinkedList) NodeContains(value int64) {
	if !(n.isEmpty()) {
		currentNode := n.getHeadNode()
		for {
			if currentNode != n.getTailNode() {
				if value == currentNode.getValue() {
					fmt.Printf("Found!!\n")
					break
				} else {
					currentNode = currentNode.getNextNode()
					continue
				}
			} else {
				break
			}
		}
	} else {
		fmt.Println("No items to the list!")
	}
}

func (n *SLinkedList) PrintList() {
	if !(n.isEmpty()) {
		currentNode := n.getHeadNode()
		for {
			if currentNode == n.getTailNode() {
				fmt.Printf("%d\n", currentNode.getValue())
				break
			} else {
				fmt.Printf("%d\n", currentNode.getValue())
				currentNode = currentNode.getNextNode()
			}
		}
	} else {
		fmt.Println("No items to the list!")
	}
}
