package main

import "fmt"

type Node struct {
	Value int
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) enqueue(value int) {
	newNode := &Node{Value: value}
	if q.tail != nil {
		q.tail.next = newNode
	}
	q.tail = newNode
	if q.head == nil {
		q.head = newNode
	}
	q.size++
}

func (q *Queue) dequeue() int {
	value := q.head.Value
	q.head = q.head.next
	q.size--
	return value
}

func (q *Queue) display() {
	temp := q.head
	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.next
	}
}

func reverse(queue *Queue) {
	if queue.head == nil || queue.head.next == nil {
		return
	}
	value := queue.dequeue()
	reverse(queue)
	queue.enqueue(value)
}

func main() {
	q := &Queue{}
	q.enqueue(30)
	q.enqueue(40)
	q.enqueue(50)
	q.display()

	reverse(q)
	q.display()
}
