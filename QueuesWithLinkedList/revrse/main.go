package main

import "fmt"

type Node struct {
	value int
	next  *Node
}
type Queue struct {
	front *Node
	rear  *Node
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
}

func (q *Queue) reverse() {
	var prev, next *Node
	current := q.front
	q.rear = q.front

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	q.front = prev
}

func (q *Queue) display() {
	temp := q.front
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}
func main() {
	q := &Queue{}
	q.enqueue(30)
	q.enqueue(440)
	q.enqueue(50)
	q.display()
	q.reverse()
	q.display()
}
