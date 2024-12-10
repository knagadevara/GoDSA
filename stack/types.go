package stack

import "cmp"

type StkElm[T cmp.Ordered] struct {
	Element   T
	FrontElem *StkElm[T]
}

type Stack[T cmp.Ordered] struct {
	Size      int
	FirstElem *StkElm[T]
}

type Queue[T cmp.Ordered] struct {
	q1, q2 Stack[T]
}

func CreateStackElem[T cmp.Ordered](Elem T) *StkElm[T] {
	return &StkElm[T]{Element: Elem, FrontElem: &StkElm[T]{}}
}
