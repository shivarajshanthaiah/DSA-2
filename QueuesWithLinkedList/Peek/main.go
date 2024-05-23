package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	front *Node
	rear  *Node
	size  int
}

func (q *Queue) enqueue(value int) {
	newNode := &Node{value: value}
	if q.rear != nil {
		q.rear.next = newNode
	}
	q.rear = newNode
	if q.front == nil {
		q.front = newNode
	}
	q.size++
}

func (q *Queue) peek() int {
	return q.front.value
}

func main() {
	queue := &Queue{}
	queue.enqueue(20)
	queue.enqueue(30)
	queue.enqueue(40)

	fmt.Println(queue.peek())
}
