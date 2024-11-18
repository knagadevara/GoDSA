package linkedlist

func Exec() {
	var sll SLinkedList
	sll.AddAtFirst(10)
	sll.AddAtLast(67)
	// sll.PrintValues()
	sll.DeleteAtFirst()
	sll.AddAtLast(51)
	// sll.PrintValues()
	sll.NodeContains(67)
	sll.AddAtFirst(1)
	// sll.PrintValues()
	sll.DeleteAtLast()
	sll.DeleteAtFirst()
	sll.DeleteAtFirst()
	sll.DeleteAtFirst()
	sll.DeleteAtFirst()
	sll.DeleteAtFirst()
	sll.DeleteAtFirst()
	sll.DeleteAtLast()
	sll.DeleteAtLast()
	sll.DeleteAtLast()
	sll.DeleteAtLast()
	sll.DeleteAtLast()
	// sll.PrintValues()
	sll.NodeContains(7)

}
