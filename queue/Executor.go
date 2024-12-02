package queue

import "fmt"

func Exec() {
	q := createQueueArr(10)
	fmt.Println(q.isEmpty())
	fmt.Println(q.isFull())
	q.EnqueueArr(5)
	q.EnqueueArr(100)
	q.EnqueueArr(45)
	q.EnqueueArr(65)
	q.EnqueueArr(75)
	q.EnqueueArr(32)
	q.PollAll()
	q.printQueStat()
	q.nextQm(35)
	q.PollAll()
	q.printQueStat()
	q.EnqueueArr(15)
	q.EnqueueArr(25)
	q.EnqueueArr(85)
	q.EnqueueArr(95)
	q.EnqueueArr(55)
	q.PollAll()
	q.printQueStat()
	q.nextQm(55)
	q.PollAll()
	q.reverseQArr()
	q.PollAll()
	q.printQueStat()
}
