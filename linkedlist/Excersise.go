package linkedlist

import "fmt"

func (n *SLinkedList[T]) swapHeadTail() {
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
		fmt.Println("Empty List")
	}
}

func (n *SLinkedList[T]) reverseList() {
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
		fmt.Println("Empty List")
	}
}

func (n *SLinkedList[T]) getKthNodeFromEnd(kth uint16) T {
	ttlElm := n.sizeOff()
	if !(n.isEmpty()) {
		currNode := n.getTailNode()
		if kth > ttlElm {
			return n.getHeadNode().getValue()
		}
		if kth == 0 {
			return n.getTailNode().getValue()
		}
		for i := 1; i <= int(kth)-1; i++ {
			currNode = currNode.getPrvNode()
		}
		return currNode.getValue()
	} else {
		fmt.Println("Empty List")
		var nilVal T
		return nilVal
	}
}
