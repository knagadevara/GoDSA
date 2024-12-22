package arrayadt

import (
	"cmp"
	"errors"
	"fmt"
)

func (a *Arrayadt[T]) IsEmpty() bool      { return a.length == 0 }
func (a *Arrayadt[T]) IsFull() bool       { return a.length == a.capasity }
func (a *Arrayadt[T]) InCap(ix int) bool  { return ix <= a.capasity-1 && ix >= 0 }
func (a *Arrayadt[T]) Len() int           { return len(a.adt) }
func (a *Arrayadt[T]) Swap(i, j int)      { a.adt[i], a.adt[j] = a.adt[j], a.adt[i] }
func (a *Arrayadt[T]) Less(i, j int) bool { return a.adt[i] < a.adt[j] }

func (a *Arrayadt[T]) Display() error {
	if !(a.IsEmpty()) {
		for i := 0; i < a.capasity; i++ {
			fmt.Println(a.adt[i])
		}
		return nil
	} else {
		return errors.New("empty-array-nothing-to-display")
	}
}

func (a *Arrayadt[T]) Addpend(val T) error {
	if !(a.IsFull()) {
		a.adt[a.length] = val
		a.length += 1
		return nil
	}
	return errors.New("capacity-full-unable-to-ppend")
}

func (a *Arrayadt[T]) Delete() error {
	if !(a.IsEmpty()) {
		var emptyVal T
		a.adt[len(a.adt)-1] = emptyVal
		a.length -= 1
		return nil
	}
	return errors.New("empty-array-nothing-to-delete")
}

func (a *Arrayadt[T]) Insert(index int, val T) error {
	if !(a.IsFull()) {
		if a.InCap(index) {
			if index > a.length-1 {
				a.adt[a.length] = val
			} else {
				for i := a.length; i >= index; i-- {
					a.adt[i+1] = a.adt[i]
				}
				a.adt[index] = val
			}
			a.length += 1
		} else {
			return errors.New("index-out-of-range")
		}
	} else {
		return errors.New("capacity-full-unable-to-nsert")
	}
	return nil
}

func (a *Arrayadt[T]) SortedInsert(val T) error {
	if !(a.IsFull()) {
		i := a.length
		for {
			if a.length == a.capasity-1 {
				return errors.New("capacity-almost-full-go-for-to-insert")
			}
			if a.adt[i] > val {
				a.adt[i+1] = a.adt[i]
				i--
				continue
			} else {
				a.adt[i+1] = val
				break
			}
		}
		a.length += 1
	} else {
		return errors.New("capacity-full-unable-to-nsert")
	}
	return nil
}

func (a *Arrayadt[T]) RmPop(index int) (T, error) {
	var nullVal T
	if !(a.IsEmpty()) {
		if a.InCap(index) {
			nullVal = a.adt[index]
			for i := index; i < a.length; i++ {
				a.adt[i] = a.adt[i+1]
			}
			a.length -= 1
			return nullVal, nil
		} else {
			return nullVal, errors.New("index-out-of-range")
		}
	}
	return nullVal, errors.New("empty-array-nothing-to-remove")
}

func (a *Arrayadt[T]) SwapI(ix1, ix2 int) error {
	if ix1 != ix2 && a.InCap(ix1) && a.InCap(ix2) {
		temp := a.adt[ix1]
		a.adt[ix1] = a.adt[ix2]
		a.adt[ix2] = temp
		return nil
	}
	return errors.New("incorrect-index-swapping")
}

func (a *Arrayadt[T]) IsSorted() bool {
	if !(a.IsEmpty()) {
		for i := 0; i < a.length-1; i++ {
			if a.adt[i] < a.adt[i+1] {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func UniqMergeSortedArray[T cmp.Ordered](A1, A2 []T) []T {
	var i, j, k, A1Len, A2Len, A3Len int
	A1Len = len(A1)
	A2Len = len(A2)
	var A3 = make([]T, (A1Len + A2Len))
	A3Len = len(A3)
	for {
		if i >= A3Len {
			break
		} else if j == A1Len && k < A2Len {
			for _, kv := range A2[k : A2Len-1] {
				A3[i] = kv
				i++
			}
			break
		} else if k == A2Len && j < A1Len {
			for _, jv := range A1[j : A1Len-1] {
				A3[i] = jv
				i++
			}
			break
		} else if j != A1Len && k != A2Len {
			if A1[j] < A2[k] {
				A3[i] = A1[j]
				if j < A1Len {
					j++
				}
			} else if A1[j] > A2[k] {
				A3[i] = A2[k]
				if k < A2Len {
					k++
				}
			} else if A1[j] == A2[k] {
				A3[i] = A2[k]
				if k < A2Len && j < A1Len {
					k++
					j++
				}
			}
			i++
		} else {
			break
		}
	}
	return A3
}

func (a *Arrayadt[T]) Set(index int, val T) error {
	if !(a.IsEmpty()) && a.InCap(index) {
		a.adt[index] = val
		return nil
	} else {
		return errors.New("index-out-of-range")
	}
}

func (a *Arrayadt[T]) Get(index int) (T, error) {
	if !(a.IsEmpty()) && a.InCap(index) {
		return a.adt[index], nil
	} else {
		var nullType T
		return nullType, errors.New("index-out-of-range")
	}
}

func (a *Arrayadt[T]) ArContains(val T) (bool, int) {
	if !(a.IsEmpty()) {
		for i := 0; i < a.capasity; i++ {
			if a.adt[i] == val {
				a.Swap(i-1, i)
				return true, i
			}
		}
	}
	return false, -1
}

func (a *Arrayadt[T]) ArContainsBS(val T) (bool, int) {
	var h, l, m int
	h = a.capasity - 1
	if !(a.IsEmpty()) {
		for {
			m = (h + l) / 2
			if val == a.adt[m] {
				return true, m
			} else {
				if val > a.adt[m] {
					l = m + 1
				} else {
					h = m - 1
				}
				if l >= h {
					break
				} else {
					continue
				}
			}
		}
	}
	return false, -1
}

func (a *Arrayadt[T]) Reverse() {
	if !(a.IsEmpty()) {
		var rivArr = make([]T, a.capasity)
		for i := a.capasity - 1; i >= 0; i-- {
			index := (a.capasity - 1) - i
			rivArr[index] = a.adt[i]
		}
		a.adt = rivArr
	}
}

func (a *Arrayadt[T]) Reverse2() {
	if !(a.IsEmpty()) {
		i := a.capasity - 1
		for {
			a.Swap((a.capasity-1)-i, i)
			i--
			if i == (a.capasity-1)/2 {
				break
			}
		}
	}
}

func (a *Arrayadt[T]) Max() T {
	var cmpOp T
	if !(a.IsEmpty()) {
		for i := 0; i <= a.capasity-1; i++ {
			if a.adt[i] > cmpOp {
				cmpOp = a.adt[i]
			}
		}
	}
	return cmpOp
}

func (a *Arrayadt[T]) Min() T {
	var cmpOp T
	if !(a.IsEmpty()) {
		for i := 0; i <= a.capasity-1; i++ {
			if a.adt[i] < cmpOp {
				cmpOp = a.adt[i]
			}
		}
	}
	return cmpOp
}
