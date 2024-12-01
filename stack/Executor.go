package stack

import "fmt"

func Exec() {
	var stk stack
	stk.Push('I')
	stk.Push('A')
	stk.Push('S')
	stk.Push('-')
	stk.Push('K')
	stk.Push('-')
	stk.Push('N')
	stk.Push('-')
	stk.Push('V')
	stk.Push('-')
	stk.PrintStack()
	stk = *stk.reverseStk()
	stk.PrintStack()
	stk.getSize()
	stk.Pop()
	stk.Pop()
	stk.Pop()
	stk.Pop()
	stk.Push('V')
	stk.PrintStack()
	stk.getSize()
	reverseString("SaiKarthik").PrintStack()
	// chkBlncdExpr("[(Sai)](VNK>)")
	fmt.Println(isBalanced("[(Sai][<VNK]>)"))
	fmt.Println(isBalanced("[(Sai)<VNK>]"))
}
