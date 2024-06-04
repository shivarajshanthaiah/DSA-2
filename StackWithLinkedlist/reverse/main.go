package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) push(value int) {
	newNode := &Node{value: value, next: s.top}
	s.top = newNode
}

func (s *Stack) display() {
	temp := s.top
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func (s *Stack) reverse() {
	var prev, next *Node
	current := s.top
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	s.top = prev
}

func main() {
	s := &Stack{}
	s.push(10)
	s.push(20)
	s.push(30)
	s.push(40)
	s.display()

	s.reverse()
	s.display()
}
