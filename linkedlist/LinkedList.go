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

func (n *SLinkedList) getNodeCount() int16 {
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
	node.setNextNode(nil)
	node.setPrvNode(nil)
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
	if n.getNodeCount() == 0 {
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
	if n.getNodeCount() == 0 {
		n.setInitialNode(node)
	} else {
		node.setPrvNode(n.getTailNode())
		n.getTailNode().setNextNode(node)
		node.setNextNode(node)
		n.setTailNode(node)
	}
	fmt.Println("Added: ", value)
	n.IncNode()
}

func (n *SLinkedList) DeleteAtFirst() {
	if n.getNodeCount() == 0 {
		fmt.Println("Nothing to Delete")
	} else {
		fmt.Printf("Deleted: %d\n", n.getHeadNode().getValue())
		n.setHeadNode(n.getHeadNode().getNextNode())
		n.DecNode()
	}
}

func (n *SLinkedList) DeleteAtLast() {
	if n.getNodeCount() == 0 {
		fmt.Println("Nothing to Delete")
	} else {
		fmt.Printf("Deleted: %d\n", n.getTailNode().getValue())
		n.setTailNode(n.getTailNode().getPrvNode())
		n.DecNode()
	}
}

func (n *SLinkedList) NodeContains(value int64) {
	headNode := n.getHeadNode()
	var indx int16
	for {
		indx += 1
		if headNode.getValue() == value {
			fmt.Printf("value %d Found at Index: %d\n", value, indx)
			break
		} else {
			if headNode.getNextNode() == headNode {
				fmt.Printf("value %d Not Found!!\n", value)
				break
			} else {
				headNode = headNode.getNextNode()
				continue
			}
		}
	}
}

func (n *SLinkedList) PrintValues() {
	headNode := n.getHeadNode()
	totalNodes := int(n.getNodeCount())
	for i := 0; i < totalNodes; i++ {
		fmt.Printf("%d\n", headNode.getValue())
		headNode = headNode.getNextNode()
	}
	fmt.Println("Total Nodes: ", totalNodes)
}
