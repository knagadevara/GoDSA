package stack

type Queue struct {
	q1, q2 stack
}

func (q *Queue) AddInQ(val rune) {
	q.q1.Push(val)
}

func (q *Queue) RemFrmQ() {
	q.q2 = *q.q1.reverseStk()
	q.q2.Pop()
	q.q1 = *q.q2.reverseStk()
}
