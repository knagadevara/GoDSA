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

func (n *SLinkedList) sizeOff() uint16 {
	return n.size
}

func (n *SLinkedList) IncElem() {
	n.size += 1
}

func (n *SLinkedList) DecElem() {
	n.size -= 1
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
	n.IncElem()
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
	n.IncElem()
	fmt.Println("Added: ", value)
}

func (n *SLinkedList) isEmpty() bool {
	return n.getHeadNode() == nil || n.getTailNode() == nil || n.sizeOff() == 0
}

func (n *SLinkedList) DeleteAtFirst() {
	if !(n.isEmpty()) {
		fmt.Printf("Deleting: %d\n", n.getHeadNode().getValue())
		secondNode := n.getHeadNode().getNextNode()
		n.getHeadNode().setPrvNode(nil)
		n.getHeadNode().setValue(0)
		n.getHeadNode().setNextNode(nil)
		n.setHeadNode(secondNode)
		n.DecElem()
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
		n.DecElem()
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
				fmt.Printf("Not Found!!\n")
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
				continue
			}
		}
	} else {
		fmt.Println("No items to the list!")
	}
}

func (n *SLinkedList) PrintSize() {
	if !(n.isEmpty()) {
		fmt.Printf("List-Size: %d\n", n.sizeOff())
	} else {
		fmt.Printf("List-Size: 0\n")
	}
}

func (n *SLinkedList) toArray() []int64 {
	if !(n.isEmpty()) {
		currentNode := n.getHeadNode()
		var llArr = make([]int64, n.sizeOff())
		for i := 0; i < int(n.size); i++ {
			llArr[i] = currentNode.getValue()
			currentNode = currentNode.getNextNode()
		}
		return llArr
	} else {
		fmt.Printf("Empty List\n")
		return nil
	}
}

func (n *SLinkedList) swapHeadTail() {
	if !(n.isEmpty()) {
		newHead := n.getTailNode()
		newHead.setNextNode(newHead.getPrvNode())
		newHead.setPrvNode(nil)
		newTail := n.getHeadNode()
		newTail.setPrvNode(newTail.getNextNode())
		newTail.setNextNode(nil)
		n.setHeadNode(newHead)
		n.setTailNode(newTail)
	} else {
		fmt.Printf("Empty List\n")
	}
}

func (n *SLinkedList) reverseList() {
	if !(n.isEmpty()) {
		n.swapHeadTail()
		currNode := n.getHeadNode().getNextNode()
		for {
			if currNode == n.getTailNode() {
				break
			} else {
				tmp := currNode.getNextNode()
				currNode.setNextNode(currNode.getPrvNode())
				currNode.setPrvNode(tmp)
				currNode = currNode.getNextNode()
				continue
			}
		}
	} else {
		fmt.Printf("Empty List\n")
	}
}

func (n *SLinkedList) getKthNodeFromEnd(kth uint16) int64 {
	ttlElm := n.sizeOff()
	if !(n.isEmpty()) {
		currNode := n.getTailNode()
		if kth >= ttlElm {
			return n.getHeadNode().getValue()
		}
		if kth == 0 {
			return n.getTailNode().getValue()
		}
		for i := 1; i <= int(kth); i++ {
			currNode = currNode.getPrvNode()
		}
		return currNode.getValue()
	} else {
		fmt.Printf("Empty List\n")
		return -1
	}
}
