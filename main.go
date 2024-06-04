package main

import "fmt"

type Node struct {
	value int
	next  *Node
}
type Stack struct {
	top  *Node
	size int
}

func (s *Stack) push(value int) {
	newNode := &Node{
		value: value,
		next:  s.top,
	}
	s.top = newNode
	s.size++
}

func (s *Stack) display() {
	temp := s.top
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func (s *Stack) pop() int {
	val := s.top.value
	s.top = s.top.next
	s.size--
	return val

}

type Queue struct {
	head *Node
	tail *Node
	size int
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
	q.size++
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func (q *Queue) display() {
	temp := q.head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func convertToQueue(stack Stack) *Queue {
	q := &Queue{}
	tempStack := &Stack{}

	for !stack.isEmpty() {
		tempStack.push(stack.pop())
	}
	for !tempStack.isEmpty() {
		q.enqueue(tempStack.pop())
	}
	return q
}
func main() {
	s := &Stack{}
	s.push(20)
	s.push(40)
	s.push(50)
	s.push(10)
	s.display()

	q := convertToQueue(*s)
	q.display()
}
