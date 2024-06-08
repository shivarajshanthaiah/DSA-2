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


func (s *Stack) popRandomNode(index int) {
	if index < 0 || index >= s.size {
		return
	}

	if index == 0 {
		s.top = s.top.next
		s.size--
		return
	}

	current := s.top
	for i := 0; i < index-1; i++ {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
	s.size--
}

func main() {
	s := &Stack{}
	s.push(30)
	s.push(40)
	s.push(50)
	s.push(60)
	s.push(70)
	s.push(70)
	s.display()

	s.popRandomNode(4)
	fmt.Println("after deleting:")
	s.display()
}
