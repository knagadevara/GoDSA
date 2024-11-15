package linkedlist

func Exec() {
	var sll SLinkedList
	sll.AddAtFirst(10)
	sll.AddAtLast(67)
	sll.PrintValues()
	sll.DeleteAtFirst()
	sll.AddAtLast(51)
	sll.PrintValues()
	sll.DeleteAtLast()
	sll.NodeContains(991)
	sll.PrintValues()
}
