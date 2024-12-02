package stack

type stkElm struct {
	element   rune
	frontElem *stkElm
}

type stack struct {
	size      int
	firstElem *stkElm
}

// const BRACKETS string = "(){}[]<>"
const R_OPEN_PRNTH rune = '('
const R_CLOSE_PRNTH rune = ')'
const R_OPEN_BRACE rune = '{'
const R_CLOSE_BRACE rune = '}'
const R_OPEN_SQUARE rune = '['
const R_CLOSE_SQUARE rune = ']'
const R_OPEN_ANGLR rune = '<'
const R_CLOSE_ANGLR rune = '>'
