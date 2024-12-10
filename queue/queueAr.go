package queue

import (
	"fmt"
)

func (q *Queue[T]) printQueStat() {
	fmt.Printf("Array Capacity:\t%d\nCurrent Token:\t%d\n", q.Cap, q.Tkn)
}

func (q *Queue[T]) isFull() bool {
	return q.Tkn == q.Cap
}

func (q *Queue[T]) isEmpty() bool {
	return q.Tkn == 0
}

func (q *Queue[T]) EnqueueArr(val T) {
	if !(q.isFull()) {
		q.ArQ[q.Tkn] = val
		q.Tkn += 1
	} else {
		fmt.Printf("Queue Full: %d\t", q.Tkn)
		fmt.Println("Dropped -> ", val)
	}
}

func (q *Queue[T]) DequeueArr() T {
	var emptyNullVal T
	if !(q.isEmpty()) {
		q.Tkn -= 1
		lftHd := q.ArQ[0]
		for i := 0; i < q.Tkn; i++ {
			q.ArQ[i] = q.ArQ[i+1]
		}
		return lftHd
	}
	return emptyNullVal
}

func (q *Queue[T]) PollAll() {
	if !(q.isEmpty()) {
		fmt.Printf("-----\n")
		for i := 0; i < q.Tkn; i++ {
			fmt.Print(" ", q.ArQ[i])
		}
		fmt.Printf("\n-----\n")
	}
}

func (q *Queue[T]) PeekHead() T {
	var emptyNullVal T
	if !(q.isEmpty()) {
		hd := q.ArQ[0]
		fmt.Println("Head :-> ", hd)
		return hd
	}
	fmt.Printf("Head :-> EMPTY\n")
	return emptyNullVal
}

func (q *Queue[T]) PeekTail() T {
	var emptyNullVal T
	if !(q.isEmpty()) {
		tl := q.ArQ[q.Tkn-1]
		fmt.Println("Tail :-> ", tl)
		return tl
	}
	fmt.Printf("Tail :-> EMPTY\n")
	return emptyNullVal
}

func (q *Queue[T]) reverseQArr() {
	if !(q.isEmpty()) {
		indx := q.Tkn - 1
		for i := 0; i <= indx/2; i++ {
			tmp := q.ArQ[i]
			q.ArQ[i] = q.ArQ[indx-i]
			q.ArQ[indx-i] = tmp
		}
	}
}

func (q *Queue[T]) nextQm(val T) {
	q.DequeueArr()
	q.EnqueueArr(val)
}

func (q *Queue[T]) PrioQueAsc(val T) {
	if !(q.isFull()) {
		q.EnqueueArr(val)
		q.sortAsc(val)
	} else {
		fmt.Printf("Queue Full: %d\t", q.Tkn)
		fmt.Println("Dropped -> ", val)
	}
}

func (q *Queue[T]) PrioQueDsc(val T) {
	if !(q.isFull()) {
		q.EnqueueArr(val)
		q.sortDsc(val)
	} else {
		fmt.Printf("Queue Full: %d\t", q.Tkn)
		fmt.Println("Dropped -> ", val)
	}
}

func (q *Queue[T]) sortAsc(val T) {
	for i := q.Tkn - 1; i >= 0; i-- {
		if val < q.ArQ[i] {
			tmp := q.ArQ[i]
			q.ArQ[i] = val
			q.ArQ[i+1] = tmp
		} else {
			continue
		}
	}
}

func (q *Queue[T]) sortDsc(val T) {
	for i := q.Tkn - 1; i > 0; i-- {
		arrVal := q.ArQ[i]
		if val >= arrVal {
			tmp := q.ArQ[i-1]
			q.ArQ[i] = tmp
			q.ArQ[i-1] = val
		} else {
			break
		}
	}
}
