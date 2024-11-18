package stack

func (e *stkElm) setElement(char rune) {
	e.element = char
}

func (e *stkElm) getElement() rune {
	return e.element
}

func (e *stkElm) setBelowElm(stkElm *stkElm) {
	e.bloElm = stkElm
}

func (e *stkElm) setAboveElm(stkElm *stkElm) {
	e.abvElm = stkElm
}

func (e *stkElm) getBelowElm() *stkElm {
	return e.bloElm
}

func (e *stkElm) getAboveElm(stkElm) *stkElm {
	return e.abvElm
}

func (s *stack) sizeOff() int {
	return s.size
}

func (s *stack) getFirstElm() *stkElm {
	return s.firstElem
}

func (s *stack) setFirstElm(firstElem *stkElm) {
	s.firstElem = firstElem
}
