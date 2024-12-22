package arrayadt

import (
	"cmp"
)

type Arrayadt[T cmp.Ordered] struct {
	adt      *[]T
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

	Concat()
	Compare()
	Copy() *Arrayadt[T]
	Union()
	Intersection()
}

func MakeArrAdt[T cmp.Ordered](sz int) *Arrayadt[T] {
	var arr = make([]T, sz)
	return &Arrayadt[T]{adt: &arr, capasity: sz, length: 0}
}
