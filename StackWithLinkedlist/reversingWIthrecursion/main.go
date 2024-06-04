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

func reverse(top *Node) *Node {
	if top == nil || top.next == nil {
		return top
	}
	newTop := reverse(top.next)
	top.next.next = top
	top.next = nil
	return newTop
}

func main() {
	s := &Stack{}
	s.push(10)
	s.push(20)
	s.push(30)
	s.push(40)
	s.display()

	rev := reverse(s.top)
	s.top = rev
	s.display()
}
