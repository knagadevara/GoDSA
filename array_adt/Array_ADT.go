package arrayadt

import (
	"errors"
	"fmt"
)

func (a *Arrayadt[T]) IsEmpty() bool {
	return a.length == 0
}

func (a *Arrayadt[T]) IsFull() bool {
	return a.length == a.capasity
}

func (a *Arrayadt[T]) InCap(ix int) bool {
	return ix <= a.capasity-1
}

func (a *Arrayadt[T]) IsBound(ix int) bool {
	return ix <= a.length-1
}

func (a *Arrayadt[T]) Display() error {
	if !(a.IsEmpty()) {
		for i := 0; i < a.length; i++ {
			fmt.Println(a.adt[i])
		}
		return nil
	} else {
		fmt.Println("Empty Array!")
		return errors.New("empty-array")
	}
}

func (a *Arrayadt[T]) Addpend(val T) error {
	if !(a.IsFull()) {
		a.adt[a.length] = val
		a.length += 1
		return nil
	}
	fmt.Println("Capacity Full!")
	return errors.New("capacity-full")
}

func (a *Arrayadt[T]) Delete() error {
	if !(a.IsEmpty()) {
		var emptyVal T
		a.adt[len(a.adt)-1] = emptyVal
		a.length -= 1
		return nil
	}
	fmt.Println("Nothing to Delete!")
	return errors.New("nothing-to-delete")
}

func (a *Arrayadt[T]) Insert(index int, val T) error {
	if !(a.IsFull()) {
		if a.InCap(index) {
			a.adt[index] = val
			if !(index < a.length) {
				a.length += 1
			}
		} else {
			fmt.Println("Index Out Of Range")
			return errors.New("index-out-of-range")
		}
	} else {
		fmt.Println("Capacity Full!")
		return errors.New("capacity-full")
	}
	return nil
}

func (a *Arrayadt[T]) Remove(index int) error {
	if !(a.IsEmpty()) && index <= a.length-1 {
		var nullVal T
		a.adt[index] = nullVal
		for i := index; i < a.length-1; i++ {
			a.adt[i] = a.adt[i+1]
		}
		a.length -= 1
		return nil
	}
	fmt.Println("Index Out Of Range")
	return errors.New("index-out-of-range")
}

func (a *Arrayadt[T]) Swap(ix1, ix2 int) error {
	if ix1 != ix2 && a.IsBound(ix1) && a.IsBound(ix2) {
		temp := a.adt[ix1]
		a.adt[ix1] = a.adt[ix2]
		a.adt[ix2] = temp
		return nil
	}
	fmt.Println("Swaping Not Possible! Check Indices.")
	return errors.New("incorrect-index-swapping")
}

func (a *Arrayadt[T]) Get(index int) (T, error) {
	if a.InCap(index) {
		return a.adt[index], nil
	} else {
		var nullType T
		return nullType, errors.New("invalid-index")
	}
}

func (a *Arrayadt[T]) ArContains(key T) (bool, int) {
	if !(a.IsEmpty()) {
		for i := 0; i < a.length; i++ {
			if key == a.adt[i] {
				return true, i
			}
		}
	}
	return false, -1
}

func (a *Arrayadt[T]) Reverse() {
	if !(a.IsEmpty()) {
		var rivArr = make([]T, a.capasity)
		for i := a.length - 1; i >= 0; i-- {
			indx := (a.length - 1) - i
			rivArr[indx] = a.adt[i]
			// fmt.Println(rivArr[indx])
		}
		a.adt = rivArr
	}
}
