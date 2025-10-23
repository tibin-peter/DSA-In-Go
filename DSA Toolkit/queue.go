package dsatoolkit

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func (q *Queue) Display() {
	fmt.Println(q.items)
}
