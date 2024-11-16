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

func (n *SLinkedList) getNodeCount() uint16 {
	return n.totalNodes
}

func (n *SLinkedList) IncNode() {
	n.totalNodes += 1
}

func (n *SLinkedList) DecNode() {
	n.totalNodes -= 1
}

func CreateNode(value int64) *IntNode {
	var node IntNode
	node.setValue(value)
	node.setNextNode(&IntNode{})
	node.setPrvNode(&IntNode{})
	return &node
}

func (n *SLinkedList) setInitialNode(node *IntNode) {
	if n.getHeadNode() == nil && n.getTailNode() == nil {
		node.setPrvNode(node)
		node.setNextNode(node)
		n.setHeadNode(node)
		n.setTailNode(node)
	}
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
	n.IncNode()
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
	n.IncNode()
}

func (n *SLinkedList) isEmpty() bool {
	return n.getHeadNode() == nil || n.getTailNode() == nil || n.getHeadNode().getNextNode() == nil || n.getTailNode().getNextNode() == nil || n.getHeadNode().getPrvNode() == nil || n.getTailNode().getPrvNode() == nil
}

func (n *SLinkedList) DeleteAtFirst() {
	if !(n.isEmpty()) {
		fmt.Printf("Deleting: %d\n", n.getHeadNode().getValue())
		n.setHeadNode(n.getHeadNode().getNextNode())
		n.getHeadNode().setPrvNode(&IntNode{})
		n.DecNode()
		if n.getHeadNode().getNextNode() == n.getHeadNode() {
			n.setHeadNode(&IntNode{})
			fmt.Println("Deleteing Last Item")
		}
	} else {
		fmt.Printf("Noting to Delete!!\n")
	}
}

func (n *SLinkedList) DeleteAtLast() {
	if !(n.isEmpty()) {
		fmt.Printf("Deleting: %d\n", n.getTailNode().getValue())
		n.setTailNode(n.getTailNode().getPrvNode())
		n.getTailNode().setNextNode(&IntNode{})
		n.DecNode()
		if n.getTailNode().getNextNode() == n.getHeadNode() {
			n.setTailNode(&IntNode{})
			fmt.Println("Deleteing Last Item")
		}
	} else {
		fmt.Printf("Noting to Delete!!\n")
	}
}

func (n *SLinkedList) NodeContains(value int64) {
	currNode := n.getHeadNode()
	var indx int16
	if !(n.isEmpty()) {
		for {
			if currNode.getValue() == value {
				fmt.Printf("value %d Found at Index: %d\n", value, indx)
				break
			} else {
				currNode = currNode.getNextNode()
				indx += 1
			}
			if currNode.getNextNode() == nil {
				break
			}
		}
	} else {
		fmt.Printf("Value: %d Not found as there are no items in Linked List!!\n", value)
	}
}

func (n *SLinkedList) PrintValues() {
	headNode := n.getHeadNode()
	totalNodes := int(n.getNodeCount())
	if !(n.isEmpty()) {
		for i := 0; i < totalNodes; i++ {
			fmt.Printf("%d\n", headNode.getValue())
			headNode = headNode.getNextNode()
		}
		fmt.Println("Total Nodes: ", totalNodes)
	} else {
		fmt.Printf("No items in Linked List!!\n")
	}
}

func (n *SLinkedList) PrintList() {
	fmt.Println(n.isEmpty())
	currentNode := n.getHeadNode()
	for {
		if currentNode == n.getTailNode() {
			break
		} else {
			fmt.Printf("%d\n", currentNode.getValue())
			currentNode = currentNode.getNextNode()
		}
	}
}
