package linkedlist

import "fmt"

func (n *Node[T]) setValue(value T) {
	n.value = value
}

func (n *Node[T]) getValue() T {
	return n.value
}

func (n *Node[T]) setNextNode(nextNode *Node[T]) {
	n.nextNode = nextNode
}

func (n *Node[T]) setPrvNode(prvNode *Node[T]) {
	n.prvNode = prvNode
}

func (n *Node[T]) getNextNode() *Node[T] {
	return n.nextNode
}

func (n *Node[T]) getPrvNode() *Node[T] {
	return n.prvNode
}

func (n *SLinkedList[T]) setHeadNode(headNode *Node[T]) {
	n.headNode = headNode
}

func (n *SLinkedList[T]) setTailNode(tailNode *Node[T]) {
	n.tailNode = tailNode
}

func (n *SLinkedList[T]) getHeadNode() *Node[T] {
	return n.headNode
}

func (n *SLinkedList[T]) getTailNode() *Node[T] {
	return n.tailNode
}

func (n *SLinkedList[T]) sizeOff() uint16 {
	return n.size
}

func (n *SLinkedList[T]) IncElem() {
	n.size += 1
}

func (n *SLinkedList[T]) DecElem() {
	n.size -= 1
}

func (n *SLinkedList[T]) isEmpty() bool {
	return n.getHeadNode() == nil || n.getTailNode() == nil || n.sizeOff() == 0
}

func (n *SLinkedList[T]) setInitialNode(node *Node[T]) {
	node.setPrvNode(node)
	node.setNextNode(node)
	n.setHeadNode(node)
	n.setTailNode(node)
}

func (n *SLinkedList[T]) AddAtFirst(value T) {
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

func (n *SLinkedList[T]) AddAtLast(value T) {
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

func (n *SLinkedList[T]) DeleteAtFirst() {
	var emptyNullVal T
	if !(n.isEmpty()) {
		fmt.Println("Deleting: ", n.getHeadNode().getValue())
		secondNode := n.getHeadNode().getNextNode()
		n.getHeadNode().setPrvNode(nil)
		n.getHeadNode().setValue(emptyNullVal)
		n.getHeadNode().setNextNode(nil)
		n.setHeadNode(secondNode)
		n.DecElem()
	} else {
		fmt.Println("Noting to Delete!!")
	}
}

func (n *SLinkedList[T]) DeleteAtLast() {
	var emptyNullVal T
	if !(n.isEmpty()) {
		fmt.Println("Deleting: ", n.getTailNode().getValue())
		lastButOne := n.getTailNode().getPrvNode()
		n.getTailNode().setNextNode(nil)
		n.getTailNode().setValue(emptyNullVal)
		n.getTailNode().setNextNode(nil)
		n.setTailNode(lastButOne)
		n.DecElem()
	} else {
		fmt.Println("Noting to Delete!!")
	}
}

func (n *SLinkedList[T]) NodeContains(value T) {
	if !(n.isEmpty()) {
		currentNode := n.getHeadNode()
		for {
			if currentNode != n.getTailNode() {
				compVal := currentNode.getValue()
				if value == compVal {
					fmt.Println("Found!!")
					break
				} else {
					currentNode = currentNode.getNextNode()
					continue
				}
			} else {
				fmt.Println("Not Found!!")
				break
			}
		}
	} else {
		fmt.Println("No items to the list!")
	}
}

func (n *SLinkedList[T]) toArray() []T {
	if !(n.isEmpty()) {
		currentNode := n.getHeadNode()
		var llArr = make([]T, n.sizeOff())
		for i := 0; i < int(n.size); i++ {
			llArr[i] = currentNode.getValue()
			currentNode = currentNode.getNextNode()
		}
		return llArr
	} else {
		fmt.Println("Empty List")
		return nil
	}
}

func (n *SLinkedList[T]) PrintList() {
	if !(n.isEmpty()) {
		currentNode := n.getHeadNode()
		for {
			if currentNode == n.getTailNode() {
				fmt.Println(currentNode.getValue())
				break
			} else {
				fmt.Println(currentNode.getValue())
				currentNode = currentNode.getNextNode()
				continue
			}
		}
	} else {
		fmt.Println("No items to the list!")
	}
}

func (n *SLinkedList[T]) PrintSize() {
	if !(n.isEmpty()) {
		fmt.Println("List-Size: ", n.sizeOff())
	} else {
		fmt.Println("List-Size: 0")
	}
}
