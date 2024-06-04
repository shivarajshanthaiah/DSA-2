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

func (q *Queue) enque(value int) {
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
		fmt.Print(temp.value," ")
		temp = temp.next
	}
	fmt.Println()
}

func (q *Queue) dequeue() int{
	frontValue := q.front.value
	q.front= q.front.next
	q.size--

	return frontValue
}

func main(){
	queue := &Queue{}

	queue.enque(30)
	queue.enque(40)
	queue.enque(10)
	queue.display()

	queue.dequeue()
	queue.display()


}
