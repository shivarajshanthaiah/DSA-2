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
	current := s.top
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}
func main() {
	stack := &Stack{}

	stack.push(25)
	stack.push(30)
	stack.push(10)

	stack.display()
}
