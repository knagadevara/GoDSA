package stack

import "fmt"

// const BRACKETS string = "(){}[]<>"
const R_OPEN_PRNTH rune = '('
const R_CLOSE_PRNTH rune = ')'
const R_OPEN_BRACE rune = '{'
const R_CLOSE_BRACE rune = '}'
const R_OPEN_SQUARE rune = '['
const R_CLOSE_SQUARE rune = ']'
const R_OPEN_ANGLR rune = '<'
const R_CLOSE_ANGLR rune = '>'

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

// // fails to check if a bracker is organically closed
// // [(Sai]<VNK]>)
// func chkBlncdExpr(word string) {
// 	var opFl, clFl int
// 	var OPEN_PRNTH stack
// 	var OPEN_BRACE stack
// 	var OPEN_SQUARE stack
// 	var OPEN_ANGLR stack
// 	var CLOSE_PRNTH stack
// 	var CLOSE_BRACE stack
// 	var CLOSE_SQUARE stack
// 	var CLOSE_ANGLR stack
// 	var ALL_LETTERS stack
// 	chkState := func() {
// 		if OPEN_PRNTH.sizeOff() != CLOSE_PRNTH.sizeOff() {
// 			if OPEN_PRNTH.sizeOff() < CLOSE_PRNTH.sizeOff() {
// 				fmt.Printf("Missing! %c\n", R_OPEN_PRNTH)
// 			} else {
// 				fmt.Printf("Missing! %c\n", R_CLOSE_PRNTH)
// 			}
// 		}
// 		if OPEN_ANGLR.sizeOff() != CLOSE_ANGLR.sizeOff() {
// 			if OPEN_ANGLR.sizeOff() < CLOSE_ANGLR.sizeOff() {
// 				fmt.Printf("Missing! %c\n", R_OPEN_ANGLR)
// 			} else {
// 				fmt.Printf("Missing! %c\n", R_CLOSE_ANGLR)
// 			}
// 		}
// 		if OPEN_BRACE.sizeOff() != CLOSE_BRACE.sizeOff() {
// 			if OPEN_BRACE.sizeOff() < CLOSE_BRACE.sizeOff() {
// 				fmt.Printf("Missing! %c\n", R_OPEN_BRACE)
// 			} else {
// 				fmt.Printf("Missing! %c\n", R_CLOSE_BRACE)
// 			}
// 		}
// 		if OPEN_SQUARE.sizeOff() != CLOSE_SQUARE.sizeOff() {
// 			if OPEN_SQUARE.sizeOff() < CLOSE_SQUARE.sizeOff() {
// 				fmt.Printf("Missing! %c\n", R_OPEN_SQUARE)
// 			} else {
// 				fmt.Printf("Missing! %c\n", R_CLOSE_SQUARE)
// 			}
// 		}
// 	}

// 	for _, v := range word {
// 		switch v {
// 		case R_OPEN_ANGLR:
// 			opFl += 1
// 			OPEN_ANGLR.Push(v)
// 		case R_CLOSE_ANGLR:
// 			clFl += 1
// 			CLOSE_ANGLR.Push(v)
// 		case R_OPEN_SQUARE:
// 			opFl += 1
// 			OPEN_SQUARE.Push(v)
// 		case R_CLOSE_SQUARE:
// 			clFl += 1
// 			CLOSE_SQUARE.Push(v)
// 		case R_OPEN_PRNTH:
// 			opFl += 1
// 			OPEN_PRNTH.Push(v)
// 		case R_CLOSE_PRNTH:
// 			clFl += 1
// 			CLOSE_PRNTH.Push(v)
// 		case R_OPEN_BRACE:
// 			opFl += 1
// 			OPEN_BRACE.Push(v)
// 		case R_CLOSE_BRACE:
// 			clFl += 1
// 			CLOSE_BRACE.Push(v)
// 		default:
// 			// clFl, opFl = false, false
// 			ALL_LETTERS.Push(v)
// 		}
// 	}
// 	// [(Sai][<VNK]>)
// 	if opFl == clFl {
// 		chkState()
// 	} else {
// 		fmt.Println("Not Balanced!")
// 	}
// }

// // [(Sai]<VNK]>)
func isBalanced(expr string) bool {
	var opFl stack
	var isBal bool
	for _, v := range expr {
		switch v {
		case R_OPEN_SQUARE, R_OPEN_PRNTH, R_OPEN_BRACE, R_OPEN_ANGLR:
			opFl.Push(v)
		case R_CLOSE_ANGLR, R_CLOSE_SQUARE, R_CLOSE_PRNTH, R_CLOSE_BRACE:
			popped := opFl.Pop().element
			if v == R_CLOSE_ANGLR && popped == R_OPEN_ANGLR {
				isBal = true
			} else if v == R_CLOSE_SQUARE && popped == R_OPEN_SQUARE {
				isBal = true
			} else if v == R_CLOSE_PRNTH && popped == R_OPEN_PRNTH {
				isBal = true
			} else if v == R_CLOSE_BRACE && popped == R_OPEN_BRACE {
				isBal = true
			} else {
				isBal = false
			}
		default:
			continue
		}
	}
	if opFl.sizeOff() == 0 && isBal {
		return true
	} else {
		return false
	}
}
