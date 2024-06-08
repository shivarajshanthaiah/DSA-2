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

func (q *Queue) display() {
	temp := q.front
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func (q *Queue) popRandom(index int) {
	if index < 0 || index >= q.size {
		return
	}
	current := q.front
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	if current.next != nil {
		current.next = current.next.next
		if current.next == nil {
			q.rear = current
		}
	}
	q.size--

}

func main() {
	q := &Queue{}
	q.enqueue(30)
	q.enqueue(40)
	q.enqueue(50)
	q.enqueue(60)
	q.display()
	q.popRandom(2)
	q.display()

}
