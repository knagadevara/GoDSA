package stack

func (q *Queue[T]) AddInQ(val T) {
	q.q1.Push(val)
}

func (q *Queue[T]) RemFrmQ() {
	q.q2 = *q.q1.reverseStk()
	q.q2.Pop()
	q.q1 = *q.q2.reverseStk()
}
