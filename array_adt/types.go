package arrayadt

import (
	"cmp"
)

type Arrayadt[T cmp.Ordered] struct {
	adt      []T
	length   int
	capasity int
}

type ArrayadtI[T cmp.Ordered] interface {
	Display() error
	IsEmpty() bool
	IsFull() bool
	IsBound(ix int) bool
	Addpend(val T) error
	Delete() error
	Insert(index int, val T) error
	Remove(index int) error
	Swap(ix1, ix2 int) error
	Get(index int) (T, error)
	ArContains(key T)
	Reverse()
}

func MakeArrAdt[T cmp.Ordered](sz int) *Arrayadt[T] {
	return &Arrayadt[T]{adt: make([]T, sz), capasity: sz}
}
