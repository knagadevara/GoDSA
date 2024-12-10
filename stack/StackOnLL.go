package stack

import "fmt"

func (e *StkElm[T]) getElement() T {
	return e.Element
}

func (e *StkElm[T]) setFrontElm(stkElm *StkElm[T]) {
	e.FrontElem = stkElm
}

func (e *StkElm[T]) getFrontElm() *StkElm[T] {
	return e.FrontElem
}

func (e *Stack[T]) sizeOff() int {
	return e.Size
}

func (e *Stack[T]) IncElem() {
	e.Size += 1
}

func (e *Stack[T]) DecElem() {
	e.Size -= 1
}

func (s *Stack[T]) getFirstElm() *StkElm[T] {
	return s.FirstElem
}

func (s *Stack[T]) setFirstElm(FirstElem *StkElm[T]) {
	s.FirstElem = FirstElem
}

func (s *Stack[T]) Push(Elem T) {
	item := CreateStackElem(Elem)
	s.IncElem()
	if s.isEmpty() {
		item.setFrontElm(&StkElm[T]{})
		s.setFirstElm(item)
	} else {
		item.setFrontElm(s.getFirstElm())
		s.setFirstElm(item)
	}
}

func (s *Stack[T]) Pop() *StkElm[T] {
	if s.isEmpty() {
		fmt.Println("Nothing to Remove!")
		return &StkElm[T]{}
	} else {
		poppedElem := s.getFirstElm()
		s.setFirstElm(s.getFirstElm().getFrontElm())
		s.DecElem()
		return poppedElem
	}
}

func (s *Stack[T]) Peek() T {
	return s.getFirstElm().getElement()
}

func (s *Stack[T]) isEmpty() bool {
	return s.sizeOff() == 0
}

func (s *Stack[T]) getSize() {
	if s.isEmpty() {
		fmt.Println("Empty Stack!")
	} else {
		fmt.Println("Size: ", s.sizeOff())
	}
}

func (s *Stack[T]) reverseStk() *Stack[T] {
	var revStk Stack[T]
	if s.isEmpty() {
		fmt.Printf("Nothing to Reverse!")
	} else {
		for {
			revStk.Push(s.Pop().Element)
			if s.isEmpty() {
				break
			}
		}
	}
	s = &revStk
	return s
}

func (s *Stack[T]) PrintStack() {
	if !(s.isEmpty()) {
		currElm := s.getFirstElm()
		for {
			fmt.Println(currElm.getElement())
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
