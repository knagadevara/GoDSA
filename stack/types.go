package stack

type stkElm struct {
	element   rune
	frontElem *stkElm
}

type stack struct {
	size      int
	firstElem *stkElm
}
