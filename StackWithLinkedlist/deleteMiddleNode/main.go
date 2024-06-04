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
	newNode := &Node{value: value, next: s.top}
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

func (s *Stack) popMiddle() {
	if s.top == nil {
		return
	}
	slow := s.top
	fast := s.top
	var prev *Node

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		prev = slow
		slow = slow.next
	}
	prev.next = slow.next

	s.size--
}

func main() {
	s := &Stack{}
	s.push(10)
	s.push(40)
	s.push(20)
	s.push(50)
	s.push(60)
	s.display()
	s.popMiddle()
	s.display()
}
