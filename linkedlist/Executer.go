package linkedlist

func Exec() {
	var sll SLinkedList
	sll.AddAtFirst(10)
	sll.AddAtLast(67)
	sll.PrintList()
	sll.DeleteAtFirst()
	sll.AddAtLast(51)
	sll.PrintList()
	sll.NodeContains(67)
	sll.AddAtFirst(1)
	sll.PrintList()
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
	sll.PrintList()
	sll.NodeContains(7)
}
