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
func main() {
	queue := &Queue{}

	queue.enque(20)
	queue.enque(30)
	queue.enque(40)

	queue.display()
}
