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

func (q *Queue) deleteMiddle() {
	if q.front == nil {
		return
	}
	slow := q.front
	fast := q.front
	var prev *Node

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		prev = slow
		slow = slow.next
	}
	prev.next = slow.next
	q.size--
}

func main() {
	q := &Queue{}
	q.enqueue(20)
	q.enqueue(30)
	q.enqueue(40)
	q.enqueue(50)
	q.enqueue(60)
	q.display()

	q.deleteMiddle()
	q.display()
}
