package queue

import (
	"fmt"
)

func createQueueArr(sz int) *QueueArr {
	return &QueueArr{ArQ: make([]int, sz), Tkn: 0, Cap: sz}
}

func (q *QueueArr) printQueStat() {
	fmt.Printf("Array Capacity:\t%d\nCurrent Token:\t%d\n", q.Cap, q.Tkn)
}

func (q *QueueArr) isFull() bool {
	return q.Tkn == q.Cap
}

func (q *QueueArr) isEmpty() bool {
	return q.Tkn == 0
}

func (q *QueueArr) EnqueueArr(val int) {
	if !(q.isFull()) {
		q.ArQ[q.Tkn] = val
		q.Tkn += 1
	}
}

func (q *QueueArr) DequeueArr() int {
	if !(q.isEmpty()) {
		q.Tkn -= 1
		lftHd := q.ArQ[0]
		for i := 0; i < q.Tkn; i++ {
			q.ArQ[i] = q.ArQ[i+1]
		}
		return lftHd
	}
	return -1
}

func (q *QueueArr) PollAll() {
	if !(q.isEmpty()) {
		fmt.Printf("[")
		for i := 0; i < q.Tkn; i++ {
			fmt.Printf(" %d ", q.ArQ[i])
		}
		fmt.Printf("]")
		fmt.Println()
	}
}

func (q *QueueArr) PeekHead() int {
	if !(q.isEmpty()) {
		hd := q.ArQ[0]
		fmt.Printf("Head :-> %d\n", hd)
		return hd
	}
	fmt.Printf("Head :-> EMPTY\n")
	return -1
}

func (q *QueueArr) PeekTail() int {
	if !(q.isEmpty()) {
		tl := q.ArQ[q.Tkn-1]
		fmt.Printf("Tail :-> %d\n", tl)
		return tl
	}
	fmt.Printf("Tail :-> EMPTY\n")
	return -1
}

func (q *QueueArr) reverseQArr() {
	if !(q.isEmpty()) {
		indx := q.Tkn - 1
		for i := 0; i <= indx/2; i++ {
			tmp := q.ArQ[i]
			q.ArQ[i] = q.ArQ[indx-i]
			q.ArQ[indx-i] = tmp
		}
	}
}

func (q *QueueArr) nextQm(val int) {
	q.DequeueArr()
	q.EnqueueArr(val)
}

func (q *QueueArr) PrioQue(val int) {
	if !(q.isFull()) {
		q.EnqueueArr(val)
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
}
