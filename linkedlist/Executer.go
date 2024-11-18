package linkedlist

import "fmt"

func Exec() {
	var sll SLinkedList
	sll.AddAtFirst(30)
	sll.AddAtFirst(20)
	sll.AddAtLast(40)
	sll.AddAtLast(50)
	sll.AddAtLast(60)
	sll.AddAtFirst(10)
	sll.AddAtLast(70)
	sll.PrintSize()
	// sll.AddAtLast(51)
	sll.PrintList()
	sll.reverseLL()
	sll.NodeContains(7)
	sll.PrintList()
	arr := sll.toArray()
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	sll.NodeContains(67)
	sll.AddAtFirst(1)
	sll.PrintList()
	sll.DeleteAtLast()
	sll.DeleteAtFirst()
	sll.DeleteAtFirst()
	sll.DeleteAtFirst()
	sll.DeleteAtLast()
	sll.DeleteAtLast()
	sll.PrintList()
	sll.PrintSize()
}
