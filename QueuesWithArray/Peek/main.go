package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) peek() int {
	return q.items[0]
}

func main() {
	queue := &Queue{}
	queue.enqueue(30)
	queue.enqueue(40)
	queue.enqueue(50)

	fmt.Println(queue.peek())
}
