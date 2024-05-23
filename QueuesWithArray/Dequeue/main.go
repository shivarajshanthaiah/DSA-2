package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) enque(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) display() {
	for _, item := range q.items {
		fmt.Println(item)
	}
}

func (q *Queue) dequeue() int {
	front := q.items[0]
	q.items = q.items[1:]
	return front
}
func main() {

	Queue := &Queue{}
	Queue.enque(20)
	Queue.enque(30)
	Queue.enque(40)
	Queue.display()

	Queue.dequeue()
	Queue.display()

}
