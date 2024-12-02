package stack

import "fmt"

func (e *stkElm) getElement() rune {
	return e.element
}

func (e *stkElm) setFrontElm(stkElm *stkElm) {
	e.frontElem = stkElm
}

func (e *stkElm) getFrontElm() *stkElm {
	return e.frontElem
}

func (e *stack) sizeOff() int {
	return e.size
}

func (e *stack) IncElem() {
	e.size += 1
}

func (e *stack) DecElem() {
	e.size -= 1
}

func (s *stack) getFirstElm() *stkElm {
	return s.firstElem
}

func (s *stack) setFirstElm(firstElem *stkElm) {
	s.firstElem = firstElem
}

func createStackElem(Elem rune) *stkElm {
	return &stkElm{element: Elem, frontElem: nil}
}

func (s *stack) Push(Elem rune) {
	item := createStackElem(Elem)
	s.IncElem()
	if s.isEmpty() {
		item.setFrontElm(nil)
		s.setFirstElm(item)
	} else {
		item.setFrontElm(s.getFirstElm())
		s.setFirstElm(item)
	}
}

func (s *stack) Pop() *stkElm {
	if s.isEmpty() {
		fmt.Printf("Nothing to Remove!\n")
		return nil
	} else {
		poppedElem := s.getFirstElm()
		s.setFirstElm(s.getFirstElm().getFrontElm())
		s.DecElem()
		return poppedElem
	}
}

func (s *stack) Peek() rune {
	return s.getFirstElm().getElement()
}

func (s *stack) isEmpty() bool {
	return s.sizeOff() == 0
}

func (s *stack) getSize() {
	if s.isEmpty() {
		fmt.Printf("Empty Stack!")
	} else {
		fmt.Printf("Size: %d\n", s.sizeOff())
	}
}

func (s *stack) PrintStack() {
	if !(s.isEmpty()) {
		currElm := s.getFirstElm()
		for {
			fmt.Printf("%c ", currElm.getElement())
			if currElm.getFrontElm() == nil {
				break
			} else {
				currElm = currElm.getFrontElm()
			}
		}
	} else {
		fmt.Printf("Empty Stack!")
	}
	fmt.Println()
}

func reverseString(word string) *stack {
	var revStk stack
	for _, v := range word {
		revStk.Push(v)
	}
	return &revStk
}

func (s *stack) reverseStk() *stack {
	var revStk stack
	fmt.Println("Size: ", s.sizeOff())
	if s.isEmpty() {
		fmt.Printf("Nothing to Reverse!")
	} else {
		for {
			revStk.Push(s.Pop().element)
			if s.isEmpty() {
				break
			}
		}
	}
	s = &revStk
	return s
}

// // [(Sai]<VNK]>)
func isBalanced(expr string) bool {
	var opFl stack
	var isBal bool
	for _, v := range expr {
		switch v {
		case R_OPEN_SQUARE, R_OPEN_PRNTH, R_OPEN_BRACE, R_OPEN_ANGLR:
			opFl.Push(v)
		case R_CLOSE_SQUARE, R_CLOSE_PRNTH, R_CLOSE_BRACE, R_CLOSE_ANGLR:
			if opFl.sizeOff() != 0 {
				popped := opFl.Pop().getElement()
				if v == R_CLOSE_ANGLR && popped == R_OPEN_ANGLR ||
					v == R_CLOSE_SQUARE && popped == R_OPEN_SQUARE ||
					v == R_CLOSE_PRNTH && popped == R_OPEN_PRNTH ||
					v == R_CLOSE_BRACE && popped == R_OPEN_BRACE {
					isBal = true
				} else {
					isBal = false
				}
			} else {
				isBal = false
			}
		default:
			continue
		}
	}
	fmt.Printf("%s: ", expr)
	if opFl.sizeOff() == 0 && isBal {
		fmt.Println("Expression is Balanced!")
		return isBal
	} else {
		fmt.Println("Expression is Not Balanced!")
		return isBal
	}
}
