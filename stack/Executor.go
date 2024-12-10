package stack

func Exec() {
	var stk Stack[rune]
	stk.Push('H')
	stk.Push('A')
	stk.Push('R')
	stk.Push('R')
	stk.Push('Y')
	stk.Push('-')
	stk.Push('J')
	stk.Push('-')
	stk.Push('P')
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
	reverseString("KaMeaLot").PrintStack()
	isBalanced("[(Harry][<Potter]>)")
	isBalanced("[(Harry)<JP>]")
}

func QueStk() {
	var qstk Queue[rune]
	qstk.RemFrmQ()
	qstk.AddInQ('R')
	qstk.AddInQ('I')
	qstk.AddInQ('D')
	qstk.AddInQ('D')
	qstk.AddInQ('L')
	qstk.AddInQ('E')

	qstk.AddInQ('T')
	qstk.AddInQ('O')
	qstk.AddInQ('M')

	qstk.RemFrmQ()
	qstk.RemFrmQ()
	qstk.RemFrmQ()

}
