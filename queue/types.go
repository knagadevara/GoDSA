package queue

import "cmp"

type QueueArr struct {
	ArQ []int
	Tkn int
	Cap int
}

type Queue[T cmp.Ordered] struct {
	ArQ []T
	Tkn int
	Cap int
}

func CreateQueueArr[T cmp.Ordered](sz int) *Queue[T] {
	return &Queue[T]{ArQ: make([]T, sz), Tkn: 0, Cap: sz}
}
