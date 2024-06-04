package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) push(value int) {
	newNode := &Node{
		value: value,
		next:  s.head,
	}
	s.head = newNode
}

func (s *Stack) pop() int {
	if s.head == nil {
		return -1
	}
	value := s.head.value
	s.head = s.head.next
	return value
}

func (s *Stack) isEmpty() bool {
	return s.head == nil
}

func (s *Stack) display() {
	temp := s.head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
	fmt.Println()
}

type Queue struct {
	head *Node
	tail *Node
}

func (q *Queue) enqueue(value int) {
	newNode := &Node{value: value}
	if q.tail != nil {
		q.tail.next = newNode
	}
	q.tail = newNode
	if q.head == nil {
		q.head = newNode
	}
}

func (q *Queue) dequeue() int {
	if q.head == nil {
		return -1
	}

	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return value
}

func (q *Queue) isEmpty() bool {
	return q.head == nil
}

func (q *Queue) display() {
	temp := q.head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
	fmt.Println()
}

func convertToStack(queue Queue) *Stack {
	stack := &Stack{}

	for !queue.isEmpty() {
		stack.push(queue.dequeue())
	}
	return stack
}

func main() {
	queue := &Queue{}

	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(30)
	queue.enqueue(60)
	queue.display()

	stack := convertToStack(*queue)
	stack.display()

	stack.pop()
	stack.display()
	fmt.Println(stack.isEmpty())
}
